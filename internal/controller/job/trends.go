package job

import (
	"lizard/internal/service"
	"lizard/internal/utils/cronjob"
)

type ITrendJobCtrl interface {
	FetchTrends(ctx *cronjob.Context)
}

func ProviderITrendsJobCtrl(trendSrv service.ITrendSrv) ITrendJobCtrl {
	return &trendJobCtrl{
		trendSrv: trendSrv,
	}
}

type trendJobCtrl struct {
	trendSrv service.ITrendSrv
}

func (t *trendJobCtrl) FetchTrends(ctx *cronjob.Context) {
	err := t.trendSrv.FetchTrends(ctx)
	if err != nil {
		return
	}
}
