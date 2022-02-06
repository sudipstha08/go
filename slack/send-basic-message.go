package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("SlackBotOAuthToken"))

	channelId, timeStamp, err := api.PostMessage("C02MRKJ65SM", slack.MsgOptionText("Hello world", false))
	if err != nil {
		fmt.Println("Error")
	}
}
