package controllers

import (
	"github.com/gin-gonic/gin"
	"lizard/source/service"
	"net/http"
)

type ITrendCtrl interface {
	FetchTrends(ctx *gin.Context)
}

func ProviderITrendsCtrl(trendSrv service.ITrendSrv) ITrendCtrl {
	return &trendCtrl{
		trendSrv: trendSrv,
	}
}

type trendCtrl struct {
	trendSrv    service.ITrendSrv
	SetResponse *StandardResponse
}

func (t *trendCtrl) FetchTrends(ctx *gin.Context) {
	err := t.trendSrv.FetchTrends(ctx)
	if err != nil {
		t.SetResponse.SetStandardResponse(ctx, http.StatusBadRequest, err)
		return
	}

	t.SetResponse.SetStandardResponse(ctx, http.StatusOK, "ok")
}
