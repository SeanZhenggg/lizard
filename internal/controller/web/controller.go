package web

import (
	"github.com/gin-gonic/gin"
	"lizard/internal/controller/web/middleware"
)

func ProvideController(lineCtrl IMessageCtrl, trendCtrl ITrendCtrl) *Controller {
	return &Controller{
		TrendCtrl:   trendCtrl,
		MessageCtrl: lineCtrl,
	}
}

type Controller struct {
	TrendCtrl   ITrendCtrl
	MessageCtrl IMessageCtrl
}

type StandardResponse struct{}

func (stdResp *StandardResponse) SetStandardResponse(ctx *gin.Context, statusCode int, data interface{}) {
	middleware.SetResp(ctx, statusCode, data)
}
