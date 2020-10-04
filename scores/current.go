package scores

import (
	"cricket/utils"
	"fmt"
)

func Current(url string) []utils.CurrentTopic {
	return utils.GetCurrentTopics(url)
}

func Commentary(url, id string) {
	comm := utils.GetCommentaryUpdates(url, id)
	var comments []string

	for _, comment := range comm.CommentaryList {
		comments = append(comments, fmt.Sprintf("🏏🏏 %v\n", comment.CommText))
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
