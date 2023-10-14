package web

import (
	"lizard/source/controllers"
	"lizard/source/controllers/middleware"

	"github.com/gin-gonic/gin"
)

type IWebApp interface {
	Init(g *gin.Engine)
}

func ProvideWebApp(
	ctrl *controllers.Controller,
	respMw middleware.IResponseMiddleware,
	authMw middleware.IAuthMiddleware,
) IWebApp {
	return &webApp{
		Ctrl:   ctrl,
		RespMw: respMw,
		AuthMw: authMw,
	}
}

type webApp struct {
	Ctrl   *controllers.Controller
	RespMw middleware.IResponseMiddleware
	AuthMw middleware.IAuthMiddleware
}

func (app *webApp) Init(g *gin.Engine) {
	app.setApiRoutes(g)
}
