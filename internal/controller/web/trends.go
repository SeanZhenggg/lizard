package web

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"lizard/internal/config"
	"lizard/internal/constant"
	"lizard/internal/model/bo"
	"lizard/internal/model/dto"
	"lizard/internal/service"
	"lizard/internal/utils/errs"
	"net/http"
	"time"
)

type ITrendCtrl interface {
	FetchTrendsAndPushMessage(ctx *gin.Context)
	RedirectToTrendPage(ctx *gin.Context)
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

func (ctrl *trendCtrl) FetchTrendsAndPushMessage(ctx *gin.Context) {
	data, err := ctrl.trendSrv.FetchTrends(ctx)
	if err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	newInserted, err := ctrl.trendSrv.UpsertTrends(ctx, data)
	if err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	messages := make([]linebot.SendingMessage, 0, len(newInserted))
	for _, r := range newInserted {
		tRes := dto.TrendResponse{
			Keyword:  r.Title,
			ShortUrl: ctrl.cfg.GetHttpConfig().BaseUrl + "/" + r.ShortUrl,
			SendTime: r.UpdatedAt.Format(time.DateTime),
		}
		messages = append(messages, linebot.NewTextMessage(tRes.Message()))
	}

	cond := &bo.SendMessage{
		To:       constant.GroupId,
		Messages: messages,
	}
	err = ctrl.messageSrv.PushMessage(ctx, cond)
	if err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	ctrl.SetResponse.SetStandardResponse(ctx, http.StatusOK, messages)
}

func (ctrl *trendCtrl) RedirectToTrendPage(ctx *gin.Context) {
	var cond dto.TrendPathCond
	err := ctx.ShouldBindUri(&cond)
	if err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, errs.CommonErr.RequestParamError)
		return
	}

	matchedTrend, err := ctrl.trendSrv.GetTrendByUrl(ctx, cond.Path)
	if err != nil {
		ctrl.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	defer ctx.Abort()
	if matchedTrend == nil {
		ctx.Redirect(http.StatusFound, constant.DefaultGoogleDailyTrendUrl)
		return
	}

	ctx.Redirect(http.StatusFound, matchedTrend.ShareUrl)
}
