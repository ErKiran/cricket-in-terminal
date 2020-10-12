package main

import (
	"cricket/scores"
	"cricket/utils"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MainCommand()
}

func MainCommand() {
	var url = os.Getenv("API_URL")
	var cmdCricket = &cobra.Command{
		Use:   "current",
		Short: "List all current Matches",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			topics := scores.Current(url)
			var matches []string
			for _, topic := range topics {
				matches = append(matches, topic.Title)
			}

			selectedMatch := ""

			prompt := &survey.Select{
				Message: "Which game you want to plug in?",
				Options: matches,
			}

			survey.AskOne(prompt, &selectedMatch)

			streamingOptions := ""
			options := &survey.Select{
				Message: "What do you want to stream on the CLI?",
				Options: []string{utils.LIVE_MATCH_COMMENTRY, utils.LIVE_SCORECARD},
			}
			survey.AskOne(options, &streamingOptions)

			var id string

			for _, topic := range topics {
				if topic.Title == selectedMatch {
					id = topic.ID
				}
			}

			if streamingOptions == utils.LIVE_MATCH_COMMENTRY {
				scores.Commentary(url, id)
			}

			if streamingOptions == utils.LIVE_SCORECARD {
				scores.QuickMatchScoreCard(url, id)
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "cricket"}
	rootCmd.AddCommand(cmdCricket)
	rootCmd.Execute()

}
