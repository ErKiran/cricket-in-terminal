package scores

import (
	"cricket/utils"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/mgutz/ansi"
	"github.com/olekukonko/tablewriter"
)

func Current(url string) []utils.CurrentTopic {
	return utils.GetCurrentTopics(url)
}

func QuickMatchScoreCard(url, id string) {
	for range time.Tick(time.Second * 20) {
		clear := make(map[string]func())

		clear["linux"] = func() {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
		value, ok := clear[runtime.GOOS]

		if ok {
			value()
		} else {
			panic("Your platform is unsupported!!")
		}

		comm := utils.GetMatchDataByID(url, id)
		if comm.Miniscore.MatchScoreDetails.State == utils.PREVIEW {
			fmt.Println(fmt.Sprintf("%s", ansi.Color("Brace Yourself Match hasn't been yet started", "red")))
			return
		}

		batTeam := comm.CommentaryList[0].BatTeamName
		var ballTeam string

		fmt.Printf("Toss:  %v won the toss and choose to %v first\n", comm.Miniscore.MatchScoreDetails.TossResults.TossWinnerName, comm.Miniscore.MatchScoreDetails.TossResults.Decision)
		fmt.Printf("%v %v\n", batTeam, comm.Miniscore.RecentOvsStats)
		fmt.Printf("%v\n", comm.Miniscore.MatchScoreDetails.CustomStatus)
		fmt.Printf("Required Run Rate %.2f\n", comm.Miniscore.RequiredRunRate)
		fmt.Printf("%v's Current Run Rate %.2f\n", batTeam, comm.Miniscore.CurrentRunRate)
		for _, inning := range comm.Miniscore.MatchScoreDetails.InningsScoreList {
			var teamName string
			if inning.BatTeamName == batTeam {
				teamName = fmt.Sprintf("%v*", inning.BatTeamName)
			} else {
				teamName = inning.BatTeamName
				ballTeam = teamName
			}
			if inning.InningsID == 1 {
				fmt.Printf("First Innings %v %d/%d(%.1f)\n\n", teamName, inning.Score, inning.Wickets, inning.Overs)
			}

			if inning.InningsID == 2 {
				fmt.Printf("Second Innings %v %d/%d(%.1f)\n\n", teamName, inning.Score, inning.Wickets, inning.Overs)
			}
		}

		fmt.Printf("Current Batting %v\n\n", batTeam)

		var strikeBatsman = comm.Miniscore.BatsmanStriker
		var nonStrikeBatsman = comm.Miniscore.BatsmanNonStriker

		batsmanData := [][]string{
			{fmt.Sprintf("%v*", strikeBatsman.BatName), fmt.Sprintf("%d", strikeBatsman.BatRuns), fmt.Sprintf("%d", strikeBatsman.BatBalls), fmt.Sprintf("%d", strikeBatsman.BatFours), fmt.Sprintf("%d", strikeBatsman.BatSixes), fmt.Sprintf("%d", strikeBatsman.BatDots), fmt.Sprintf("%.2f", strikeBatsman.BatStrikeRate)},
			{nonStrikeBatsman.BatName, fmt.Sprintf("%d", nonStrikeBatsman.BatRuns), fmt.Sprintf("%d", nonStrikeBatsman.BatBalls), fmt.Sprintf("%d", nonStrikeBatsman.BatFours), fmt.Sprintf("%d", nonStrikeBatsman.BatSixes), fmt.Sprintf("%d", nonStrikeBatsman.BatDots), fmt.Sprintf("%.2f", nonStrikeBatsman.BatStrikeRate)},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"NAME", "RUNS", "BALLS", "FOURS", "SIXS", "DOTS", "S/R"})
		for _, v := range batsmanData {
			table.Append(v)
		}
		table.Render()
		fmt.Printf("\n\n\n")

		fmt.Printf("Current Bowling %v\n\n", ballTeam)

		var strikeBowler = comm.Miniscore.BowlerStriker
		var nonStrikeBowler = comm.Miniscore.BowlerNonStriker

		bowlingTable := tablewriter.NewWriter(os.Stdout)
		bowlingData := [][]string{
			{fmt.Sprintf("%v*", strikeBowler.BowlName), fmt.Sprintf("%d", strikeBowler.BowlWkts), fmt.Sprintf("%.1f", strikeBowler.BowlOvs), fmt.Sprintf("%d", strikeBowler.BowlWides), fmt.Sprintf("%d", strikeBowler.BowlNoballs), fmt.Sprintf("%d", strikeBowler.BowlMaidens), fmt.Sprintf("%.2f", strikeBowler.BowlEcon)},
			{nonStrikeBowler.BowlName, fmt.Sprintf("%d", nonStrikeBowler.BowlWkts), fmt.Sprintf("%.1f", nonStrikeBowler.BowlOvs), fmt.Sprintf("%d", nonStrikeBowler.BowlWides), fmt.Sprintf("%d", nonStrikeBowler.BowlNoballs), fmt.Sprintf("%d", nonStrikeBowler.BowlMaidens), fmt.Sprintf("%.2f", nonStrikeBowler.BowlEcon)},
		}
		bowlingTable.SetHeader([]string{"NAME", "WICKETS", "OVERS", "WIDES", "NOBALLS", "MAIDENS", "ECONOMY"})
		for _, v := range bowlingData {
			bowlingTable.Append(v)
		}
		bowlingTable.Render()
	}
}

func Commentary(url, id string) {
	// for range time.Tick(time.Second * 1) {
	comm := utils.GetMatchDataByID(url, id)
	var comments []string

	for _, comment := range comm.CommentaryList {
		event := comment.Event
		var emoji string
		if event == "NONE" {
			event = ""
		}

		if event == utils.FOUR || event == utils.SIX {
			emoji = "🏏"
		}

		if event == utils.OVER_BREAK {
			emoji = "⏳"
		}

		if event == utils.WICKET {
			emoji = "🇼"
		}

		if len(comment.CommentaryFormats.Bold.FormatValue) != 0 {
			comments = append(comments, fmt.Sprintf("%v  %v %v\n", emoji, event, strings.Replace(comment.CommText, comment.CommentaryFormats.Bold.FormatID[0], comment.CommentaryFormats.Bold.FormatValue[0], -1)))
		} else {
			comments = append(comments, fmt.Sprintf("%v  %v %v\n", emoji, event, comment.CommText))
		}

		// batsman := fmt.Sprintf(`%v %v (%v) [(%d) 6's, (%d) 4's]`, comment.BatsmanStriker.BatName, comment.BatsmanStriker.BatRuns, comment.BatsmanStriker.BatBalls, comment.BatsmanStriker.BatSixes, comment.BatsmanStriker.BatFours)
		// bowler := fmt.Sprintf(`%v %v-%v(%.1f)`, comment.BowlerStriker.BowlName, comment.BowlerStriker.BowlWkts, comment.BowlerStriker.BowlRuns, comment.BowlerStriker.BowlOvs)

		// for _, innings := range comment.{
		// 	fmt.Println(innings)
		// }

		// fmt.Println(batsman)
		// fmt.Println(bowler)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
	// }
}
