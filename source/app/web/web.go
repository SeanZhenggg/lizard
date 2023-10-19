package web

import (
	"lizard/source/controllers/web"
	middleware2 "lizard/source/controllers/web/middleware"

	"github.com/gin-gonic/gin"
)

type IWebApp interface {
	Init(g *gin.Engine)
}

func ProvideWebApp(
	ctrl *web.Controller,
	respMw middleware2.IResponseMiddleware,
	authMw middleware2.IAuthMiddleware,
) IWebApp {
	return &webApp{
		Ctrl:   ctrl,
		RespMw: respMw,
		AuthMw: authMw,
	}
}

type webApp struct {
	Ctrl   *web.Controller
	RespMw middleware2.IResponseMiddleware
	AuthMw middleware2.IAuthMiddleware
}

func (app *webApp) Init(g *gin.Engine) {
	app.setApiRoutes(g)
}
