package cronjob

import (
	"context"
	"github.com/SeanZhenggg/go-utils/logger"
	"github.com/robfig/cron/v3"
)

type ICronJob interface {
	Use(job ...FuncJob)
	Start()
	Stop() context.Context
	AddScheduleFunc(spec string, cmd FuncJob) (cron.EntryID, error)
}

func ProviderCronJob(logger logger.ILogger) ICronJob {
	return &cronJob{
		cron:   cron.New(cron.WithSeconds()),
		logger: logger,
	}
}

type cronJob struct {
	cron      *cron.Cron
	handlers  handlerChain
	logger    logger.ILogger
	closeChan chan struct{}
}

func (cj *cronJob) AddScheduleFunc(spec string, cmd FuncJob) (cron.EntryID, error) {
	return cj.addScheduleFunc(spec, cmd)
}

func (cj *cronJob) Use(job ...FuncJob) {
	cj.handlers = append(cj.handlers, job...)
}

func (cj *cronJob) Start() {
	cj.cron.Start()
}

func (cj *cronJob) Stop() context.Context {
	return cj.cron.Stop()
}

func (cj *cronJob) addScheduleFunc(spec string, cmd FuncJob) (cron.EntryID, error) {
	job := NewCustomJobFunc(cj, cmd)
	return cj.cron.AddJob(spec, job)
}
