package bo

import "github.com/line/line-bot-sdk-go/v7/linebot"

type SendMessage struct {
	To       string
	Messages []linebot.SendingMessage
}
