package job

import (
	"context"
	"github.com/SeanZhenggg/go-utils/logger"
	"golang.org/x/xerrors"
	"lizard/internal/controller/job"
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
) IJobApp {
	return &jobApp{
		ctrl:   ctrl,
		cron:   cron,
		logger: logger,
	}
}

type jobApp struct {
	ctrl   *job.Controller
	cron   cronjob.ICronJob
	logger logger.ILogger
}

func (app *jobApp) Init() {
	app.cron.Use(func(ctx *cronjob.Context) {
		app.logger.Info("this is middleware start...")
		ctx.Next()
		app.logger.Info("this is middleware end...")
	})

	//_, err := app.cron.AddScheduleFunc("*/3 * * * * *", func(ctx *cronjob.Context) {
	//	app.logger.Info("running job do something...")
	//	<-time.After(5 * time.Second)
	//	app.logger.Info("running job completed...")
	//})
	_, err := app.cron.AddScheduleFunc("*/3 * * * * *", app.ctrl.TrendJobCtrl.SendToGroup)

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
