package job

import (
	"lizard/source/controllers/web"
	"lizard/source/controllers/web/middleware"
)

type IJobApp interface {
	Init()
}

func ProvideJobApp(
	ctrl *web.Controller,
	respMw middleware.IResponseMiddleware,
	authMw middleware.IAuthMiddleware,
) IJobApp {
	return &jobApp{
		Ctrl:   ctrl,
		RespMw: respMw,
		AuthMw: authMw,
	}
}

type jobApp struct {
	Ctrl   *web.Controller
	RespMw middleware.IResponseMiddleware
	AuthMw middleware.IAuthMiddleware
}

func (app *jobApp) Init() {

}
