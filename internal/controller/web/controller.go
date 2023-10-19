package web

import (
	"github.com/gin-gonic/gin"
	"lizard/internal/controller/web/middleware"
)

func ProvideController(trendCtrl ITrendCtrl) *Controller {
	return &Controller{
		TrendCtrl: trendCtrl,
	}
}

type Controller struct {
	TrendCtrl ITrendCtrl
}

type StandardResponse struct{}

func (stdResp *StandardResponse) SetStandardResponse(ctx *gin.Context, statusCode int, data interface{}) {
	middleware.SetResp(ctx, statusCode, data)
}