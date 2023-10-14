package controllers

import (
	"github.com/gin-gonic/gin"
	"lizard/source/controllers/middleware"
)

func ProvideController(trendCtrl ITrendCtrl) *Controller {
	return &Controller{
		trendCtrl: trendCtrl,
	}
}

type Controller struct {
	trendCtrl ITrendCtrl
}

type StandardResponse struct{}

func (stdResp *StandardResponse) SetStandardResponse(ctx *gin.Context, statusCode int, data interface{}) {
	middleware.SetResp(ctx, statusCode, data)
}
