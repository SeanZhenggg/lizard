package web

import (
	"lizard/internal/controller/web"
	"lizard/internal/controller/web/middleware"

	"github.com/gin-gonic/gin"
)

type IWebApp interface {
	Init(g *gin.Engine)
}

func ProvideWebApp(
	ctrl *web.Controller,
	respMw middleware.IResponseMiddleware,
) IWebApp {
	return &webApp{
		Ctrl:   ctrl,
		RespMw: respMw,
	}
}

type webApp struct {
	Ctrl   *web.Controller
	RespMw middleware.IResponseMiddleware
}

func (app *webApp) Init(g *gin.Engine) {
	app.setApiRoutes(g)
}
