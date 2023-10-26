package job

import (
	"context"
	"github.com/SeanZhenggg/go-utils/logger"
	"golang.org/x/xerrors"
	"lizard/internal/controller/job"
	"lizard/internal/controller/job/middleware"
	"lizard/internal/utils/cronjob"
)

type IJobApp interface {
	Init()
	Start()
	Stop() context.Context
}

func ProvideJobApp(
	ctrl *job.Controller,
	cron cronjob.ICronJob,
	logger logger.ILogger,
	mw middleware.IJobLogMiddleware,
) IJobApp {
	return &jobApp{
		ctrl:   ctrl,
		mw:     mw,
		cron:   cron,
		logger: logger,
	}
}

type jobApp struct {
	ctrl   *job.Controller
	cron   cronjob.ICronJob
	logger logger.ILogger
	mw     middleware.IJobLogMiddleware
}

func (app *jobApp) Init() {
	app.cron.Use(app.mw.Handle)

	_, err := app.cron.AddScheduleFunc("* */1 * * * *", app.ctrl.TrendJobCtrl.FetchTrendsAndPushMessage)

	if err != nil {
		app.logger.Error(xerrors.Errorf("jobApp Init app.cron.AddScheduleFunc error: %w", err))
		return
	}
}

func (app *jobApp) Start() {
	app.cron.Start()
}

func (app *jobApp) Stop() context.Context {
	return app.cron.Stop()
}
