package web

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"lizard/internal/config"
	"lizard/internal/constant"
	"lizard/internal/model/bo"
	"lizard/internal/model/dto"
	"lizard/internal/service"
	"net/http"
	"time"
)

type ITrendCtrl interface {
	FetchTrends(ctx *gin.Context)
}

func ProviderITrendsCtrl(trendSrv service.ITrendSrv, messageSrv service.IMessageSrv, cfg config.IConfigEnv) ITrendCtrl {
	return &trendCtrl{
		trendSrv:   trendSrv,
		messageSrv: messageSrv,
		cfg:        cfg,
	}
}

type trendCtrl struct {
	trendSrv    service.ITrendSrv
	messageSrv  service.IMessageSrv
	cfg         config.IConfigEnv
	SetResponse *StandardResponse
}

func (t *trendCtrl) FetchTrends(ctx *gin.Context) {
	data, err := t.trendSrv.FetchTrends(ctx)
	if err != nil {
		t.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	newInserted, err := t.trendSrv.UpsertTrends(ctx, data)
	if err != nil {
		t.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
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
		t.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	t.SetResponse.SetStandardResponse(ctx, http.StatusOK, messages)
}
