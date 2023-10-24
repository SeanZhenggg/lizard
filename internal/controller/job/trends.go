package job

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"lizard/internal/constant"
	"lizard/internal/model/bo"
	"lizard/internal/service"
	"lizard/internal/utils/cronjob"
	"log"
)

type ITrendJobCtrl interface {
	FetchTrends(ctx *cronjob.Context)
	SendToGroup(ctx *cronjob.Context)
}

func ProviderITrendsJobCtrl(messageSrv service.IMessageSrv, trendSrv service.ITrendSrv) ITrendJobCtrl {
	return &trendJobCtrl{
		trendSrv:   trendSrv,
		messageSrv: messageSrv,
	}
}

type trendJobCtrl struct {
	trendSrv   service.ITrendSrv
	messageSrv service.IMessageSrv
}

func (t *trendJobCtrl) FetchTrends(ctx *cronjob.Context) {
	err := t.trendSrv.FetchTrends(ctx)
	if err != nil {
		return
	}
}

func (t *trendJobCtrl) SendToGroup(ctx *cronjob.Context) {
	cond := &bo.SendMessage{To: constant.GroupId, Messages: []linebot.SendingMessage{linebot.NewTextMessage("hello!!!")}}
	err := t.messageSrv.PushMessage(ctx, cond)
	if err != nil {
		log.Printf("err : %v\n", err)
		return
	}
}
