package dto

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type TrendResponse struct {
	Keyword  string
	ShortUrl string
	SendTime string
}

func (tr TrendResponse) Message() string {
	return fmt.Sprintf("關鍵字：%s\n\n網址：%s\n\n%s", tr.Keyword, tr.ShortUrl, tr.SendTime)
}

type SendMessage struct {
	To       string
	Messages []linebot.SendingMessage
}

type LineResponse struct {
	Raw     string `json:"raw"`
	Source  string `json:"source"`
	Message string `json:"message"`
}
