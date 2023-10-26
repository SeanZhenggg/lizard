package job

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"lizard/internal/config"
	"lizard/internal/constant"
	"lizard/internal/model/bo"
	"lizard/internal/model/dto"
	"lizard/internal/service"
	"lizard/internal/utils/cronjob"
	"time"
)

type ITrendJobCtrl interface {
	FetchTrendsAndPushMessage(ctx *cronjob.Context)
}

func ProviderITrendsJobCtrl(messageSrv service.IMessageSrv, trendSrv service.ITrendSrv, cfg config.IConfigEnv) ITrendJobCtrl {
	return &trendJobCtrl{
		trendSrv:   trendSrv,
		messageSrv: messageSrv,
		cfg:        cfg,
	}
}

type trendJobCtrl struct {
	trendSrv   service.ITrendSrv
	messageSrv service.IMessageSrv
	cfg        config.IConfigEnv
}

func (t *trendJobCtrl) FetchTrendsAndPushMessage(ctx *cronjob.Context) {
	defer SetJobFunc(ctx)

	data, err := t.trendSrv.FetchTrends(ctx)
	if err != nil {
		SetJobError(ctx, err)
		return
	}

	newInserted, err := t.trendSrv.UpsertTrends(ctx, data)
	if err != nil {
		SetJobError(ctx, err)
		return
	}

	messages := make([]linebot.SendingMessage, 0, len(newInserted))
	for _, r := range newInserted {
		tRes := dto.TrendResponse{
			Keyword:  r.Title,
			ShortUrl: t.cfg.GetHttpConfig().BaseUrl + "/" + r.ShortUrl,
			SendTime: r.UpdatedAt.Format(time.DateTime),
		}
		messages = append(messages, linebot.NewTextMessage(tRes.Message()))
	}

	cond := &bo.SendMessage{
		To:       constant.GroupId,
		Messages: messages,
	}
	err = t.messageSrv.PushMessage(ctx, cond)
	if err != nil {
		SetJobError(ctx, err)
		return
	}
}
