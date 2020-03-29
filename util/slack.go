package util

import (
	"github.com/slack-go/slack"
)

const (
	COLOR_SUCCESS = "#4dd14d"
	COLOR_FAIL    = "#FF0000"
)

var (
	COLOR_SUCCESS_LIST = []string{"#22f98d", "#00b95d", "#368992"}
)

func SendAttatchment(client *slack.Client, channel string, text string, attachments ...slack.Attachment) {
	client.PostMessage(channel, 
		slack.MsgOptionText(text, false), 
		slack.MsgOptionAttachments(attachments...),
	)
}

func SendError(client *slack.Client, channel string, err error) {
	client.PostMessage(channel, 
		slack.MsgOptionUsername("awsbot"),	
		slack.MsgOptionAttachments(
			slack.Attachment{
				Text:       err.Error(),
				Color:      COLOR_FAIL,
				MarkdownIn: []string{"text"},
			},
		),
	)
}
