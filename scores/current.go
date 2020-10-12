package scores

import (
	"cricket/utils"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func Current(url string) []utils.CurrentTopic {
	return utils.GetCurrentTopics(url)
}

func QuickMatchScoreCard(url, id string) {
	comm := utils.GetMatchDataByID(url, id)

	if comm.Miniscore.MatchScoreDetails.State == utils.PREVIEW {
		fmt.Println("Brace Yourself Match hasn't been yet started")
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
			fmt.Printf("First Innings %v %d/%d(%.1f)\n", teamName, inning.Score, inning.Wickets, inning.Overs)
		}

		if inning.InningsID == 2 {
			fmt.Printf("Second Innings %v %d/%d(%.1f)\n", teamName, inning.Score, inning.Wickets, inning.Overs)
		}
	}

	fmt.Printf("Current Batting %v\n", batTeam)

	var strikeBatsman = comm.Miniscore.BatsmanStriker
	var nonStrikeBatsman = comm.Miniscore.BatsmanNonStriker

	batsmanData := [][]string{
		{fmt.Sprintf("%v*", strikeBatsman.BatName), fmt.Sprintf("%d", strikeBatsman.BatRuns), fmt.Sprintf("%d", strikeBatsman.BatBalls), fmt.Sprintf("%d", strikeBatsman.BatFours), fmt.Sprintf("%d", strikeBatsman.BatSixes), fmt.Sprintf("%d", strikeBatsman.BatDots), fmt.Sprintf("%.2f", strikeBatsman.BatStrikeRate)},
		{nonStrikeBatsman.BatName, fmt.Sprintf("%d", nonStrikeBatsman.BatRuns), fmt.Sprintf("%d", nonStrikeBatsman.BatBalls), fmt.Sprintf("%d", nonStrikeBatsman.BatFours), fmt.Sprintf("%d", nonStrikeBatsman.BatSixes), fmt.Sprintf("%d", nonStrikeBatsman.BatDots), fmt.Sprintf("%.2f", nonStrikeBatsman.BatStrikeRate)},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetCaption(true, "BATING SCORECARD.")
	table.SetHeader([]string{"NAME", "RUNS", "BALLS", "FOURS", "SIXS", "DOTS", "S/R"})
	for _, v := range batsmanData {
		table.Append(v)
	}
	table.Render()

	fmt.Printf("Current Bowling %v\n", ballTeam)

	var strikeBowler = comm.Miniscore.BowlerStriker
	var nonStrikeBowler = comm.Miniscore.BowlerNonStriker

	fmt.Printf("%v* %v-%v(%.1f)\n", strikeBowler.BowlName, strikeBowler.BowlWkts, strikeBowler.BowlRuns, strikeBowler.BowlOvs)
	fmt.Printf("%v %v-%v(%.1f)\n", nonStrikeBowler.BowlName, nonStrikeBowler.BowlWkts, nonStrikeBowler.BowlRuns, nonStrikeBowler.BowlOvs)

	// data := [][]string{
	// 	{"A", "The Good", "500"},
	// 	{"B", "The Very very Bad Man", "288"},
	// 	{"C", "The Ugly", "120"},
	// 	{"D", "The Gopher", "800"},
	// }

	// table := tablewriter.NewWriter(os.Stdout)
	// table.SetHeader([]string{"Name", "Sign", "Rating"})
	// table.SetCaption(true, "Movie ratings.")

	// for _, v := range data {
	// 	table.Append(v)
	// }
	// table.Render()
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
