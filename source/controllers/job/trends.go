package job

import (
	"github.com/gin-gonic/gin"
	"lizard/source/service"
)

type ITrendCtrl interface {
	FetchTrends(ctx *gin.Context)
}

func ProviderITrendsJobCtrl(trendSrv service.ITrendSrv) ITrendCtrl {
	return &trendJobCtrl{
		trendSrv: trendSrv,
	}
}

type trendJobCtrl struct {
	trendSrv service.ITrendSrv
}

func (t *trendJobCtrl) FetchTrends(ctx *gin.Context) {

}
