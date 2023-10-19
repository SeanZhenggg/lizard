package job

import (
	"lizard/source/controllers/web"
	middleware2 "lizard/source/controllers/web/middleware"
)

type IJobApp interface {
	Init()
}

func ProvideJobApp(
	ctrl *web.Controller,
	respMw middleware2.IResponseMiddleware,
	authMw middleware2.IAuthMiddleware,
) IJobApp {
	return &jobApp{
		Ctrl:   ctrl,
		RespMw: respMw,
		AuthMw: authMw,
	}
}

type jobApp struct {
	Ctrl   *web.Controller
	RespMw middleware2.IResponseMiddleware
	AuthMw middleware2.IAuthMiddleware
}

func (app *jobApp) Init() {

}
