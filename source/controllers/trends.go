package controllers

import (
	"github.com/gin-gonic/gin"
	"lizard/source/mongo"
	"lizard/source/service"
)

type ITrendCtrl interface {
	GetTrends(ctx *gin.Context)
}

func ProviderITrendsCtrl(trendSrv service.ITrendSrv, db mongo.IMongoCli) ITrendCtrl {
	return &trendCtrl{
		db:       db,
		trendSrv: trendSrv,
	}
}

type trendCtrl struct {
	db       mongo.IMongoCli
	trendSrv service.ITrendSrv
}

func (t *trendCtrl) GetTrends(ctx *gin.Context) {
	t.trendSrv.GetTrends(ctx)
}
