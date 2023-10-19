package job

import (
	"github.com/gin-gonic/gin"
	"lizard/internal/service"
)

type ITrendJobCtrl interface {
	FetchTrends(ctx *gin.Context)
}

func ProviderITrendsJobCtrl(trendSrv service.ITrendSrv) ITrendJobCtrl {
	return &trendJobCtrl{
		trendSrv: trendSrv,
	}
}

type trendJobCtrl struct {
	trendSrv service.ITrendSrv
}

func (t *trendJobCtrl) FetchTrends(ctx *gin.Context) {
	err := t.trendSrv.FetchTrends(ctx)
	if err != nil {
		return
	}
}
