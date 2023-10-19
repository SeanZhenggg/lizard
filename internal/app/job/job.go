package job

import (
	"github.com/SeanZhenggg/go-utils/logger"
	"github.com/robfig/cron/v3"
	"lizard/internal/controller/job"
)

type IJobApp interface {
	Start()
	Stop()
}

func ProvideJobApp(
	ctrl *job.Controller,
	logger logger.ILogger,
) IJobApp {
	return &jobApp{
		Ctrl:   ctrl,
		logger: logger,
	}
}

type jobApp struct {
	Ctrl   *job.Controller
	cron   *cron.Cron
	logger logger.ILogger
}

func (app *jobApp) Init() {
	app.cron = cron.New(cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))
}

func (app *jobApp) Start() {

}

func (app *jobApp) Stop() {

}
