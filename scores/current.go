package scores

import (
	"cricket/utils"
	"fmt"
	"strings"
)

func Current(url string) []utils.CurrentTopic {
	return utils.GetCurrentTopics(url)
}

func Commentary(url, id string) {
	comm := utils.GetCommentaryUpdates(url, id)
	var comments []string

	for _, comment := range comm.CommentaryList {
		if len(comment.CommentaryFormats.Bold.FormatValue) != 0 {
			comments = append(comments, fmt.Sprintf("🏏🏏 %v\n", strings.Replace(comment.CommText, comment.CommentaryFormats.Bold.FormatID[0], comment.CommentaryFormats.Bold.FormatValue[0], -1)))
		} else {
			comments = append(comments, fmt.Sprintf("🏏🏏 %v\n", comment.CommText))
		}
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
