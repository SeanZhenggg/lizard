package job

import (
	"lizard/internal/controller/job"
)

type IJobApp interface {
	Init()
}

func ProvideJobApp(
	ctrl *job.Controller,
) IJobApp {
	return &jobApp{
		Ctrl: ctrl,
	}
}

type jobApp struct {
	Ctrl *job.Controller
}

func (app *jobApp) Init() {

}
