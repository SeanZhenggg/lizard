package job

import (
	"github.com/SeanZhenggg/go-utils/logger"
	"github.com/gin-gonic/gin"
	"lizard/internal/service"
)

type ITrendCtrl interface {
	FetchTrends(ctx *gin.Context)
}

func ProviderITrendsJobCtrl(trendSrv service.ITrendSrv, logger logger.ILogger) ITrendCtrl {
	return &trendJobCtrl{
		trendSrv: trendSrv,
		logger:   logger,
	}
}

type trendJobCtrl struct {
	trendSrv service.ITrendSrv
	logger   logger.ILogger
}

func (t *trendJobCtrl) FetchTrends(ctx *gin.Context) {
	err := t.trendSrv.FetchTrends(ctx)
	if err != nil {
		return
	}
}
