package scores

import (
	"cricket/utils"
	"fmt"
	"strings"
)

func Current(url string) []utils.CurrentTopic {
	return utils.GetCurrentTopics(url)
}

func QuickMatchScoreCard(url, id string) {
	comm := utils.GetMatchDataByID(url, id)

	for _, inning := range comm.Miniscore.MatchScoreDetails.InningsScoreList {
		if inning.InningsID == 1 {
			fmt.Printf("First Innings %v %d-%d(%.1f)", inning.BatTeamName, inning.Score, inning.Wickets, inning.Overs)
		}
	}

	fmt.Println(comm.Miniscore.MatchScoreDetails.CustomStatus)
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
