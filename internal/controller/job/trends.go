package job

import (
	"lizard/internal/constant"
	"lizard/internal/model/bo"
	"lizard/internal/service"
	"lizard/internal/utils/cronjob"
	"log"
)

type ITrendJobCtrl interface {
	FetchTrendsAndPush(ctx *cronjob.Context)
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

func (t *trendJobCtrl) FetchTrendsAndPush(ctx *cronjob.Context) {
	err := t.trendSrv.FetchTrends(ctx)
	if err != nil {
		// TODO: send log to cron middleware
		return
	}

	cond := &bo.SendMessage{
		To: constant.GroupId,
		//Messages:
	}
	err = t.messageSrv.PushMessage(ctx, cond)
	if err != nil {
		// TODO: send log to cron middleware
		log.Printf("err : %v\n", err)
		return
	}
}

func (t *trendJobCtrl) SendToGroup(ctx *cronjob.Context, cond *bo.SendMessage) {
	err := t.messageSrv.PushMessage(ctx, cond)
	if err != nil {
		log.Printf("err : %v\n", err)
		return
	}
}
