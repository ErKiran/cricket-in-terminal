package scores

import (
	"cricket/utils"
)

func Current(url string) []utils.CurrentTopic {
	return utils.GetCurrentTopics(url)
}

func Commentary(url, id string) {
	utils.GetCommentaryUpdates(url, id)
}
