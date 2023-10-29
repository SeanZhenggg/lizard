package web

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"lizard/internal/constant"
	"log"
)

type IMessageCtrl interface {
	RecvMessage(ctx *gin.Context)
}

func ProvideMessageCtrl() IMessageCtrl {
	return &messageCtrl{}
}

type messageCtrl struct {
	SetResponse *StandardResponse
}

func (ctrl *messageCtrl) RecvMessage(ctx *gin.Context) {
	bot, err := linebot.New(constant.ChannelSecret, constant.ChannelAccessToken)
	if err != nil {
		log.Printf("messageCtrl bot.ParseRequest: %v\n", err)
		return
	}

	events, err := bot.ParseRequest(ctx.Request)
	if err != nil {
		log.Printf("messageCtrl bot.ParseRequest: %v\n", err)
		return
	}

	for _, event := range events {
		log.Printf("event : %+v\n", event)
		log.Printf("event source: %+v\n", event.Source)
		log.Printf("event message: %+v\n", event.Message)
	}
}
