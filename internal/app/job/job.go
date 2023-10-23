package job

import (
	"github.com/SeanZhenggg/go-utils/logger"
	"golang.org/x/xerrors"
	"lizard/internal/controller/job"
	"lizard/internal/utils/cronjob"
	"log"
)

type IJobApp interface {
	Init()
	Start()
	Stop()
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
		log.Printf("this is middleware start...")
		ctx.Next()
		log.Printf("this is middleware end...")
	})

	//_, err := app.cron.AddScheduleFunc("* * * * *", app.ctrl.TrendJobCtrl.FetchTrends)
	_, err := app.cron.AddScheduleFunc("* * * * *", func(ctx *cronjob.Context) {
		log.Printf("running job do something...")
	})

	if err != nil {
		app.logger.Error(xerrors.Errorf("jobApp Init error: %w", err))
		return
	}
}

func (app *jobApp) Start() {
	log.Printf("jobApp Start")
	app.cron.Start()
}

func (app *jobApp) Stop() {
	app.cron.Stop()
}
