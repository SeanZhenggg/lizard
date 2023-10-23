package cronjob

import (
	"github.com/SeanZhenggg/go-utils/logger"
	"github.com/robfig/cron/v3"
	"log"
)

type ICronJob interface {
	Use(job ...FuncJob)
	Start()
	Stop()
	AddScheduleFunc(spec string, cmd FuncJob) (cron.EntryID, error)
}

func ProviderCronJob(logger logger.ILogger) ICronJob {
	return &cronJob{
		cron:   cron.New(),
		logger: logger,
	}
}

type cronJob struct {
	cron     *cron.Cron
	handlers handlerChain
	logger   logger.ILogger
}

func (cj *cronJob) AddScheduleFunc(spec string, cmd FuncJob) (cron.EntryID, error) {
	return cj.addScheduleFunc(spec, cmd)
}

func (cj *cronJob) Use(job ...FuncJob) {
	cj.handlers = append(cj.handlers, job...)
}

func (cj *cronJob) Start() {
	log.Printf("cronjob Start")
	cj.cron.Start()
}

func (cj *cronJob) Stop() {
	ctx := cj.cron.Stop()

	log.Printf("running cron job not completed...")
	<-ctx.Done()
	log.Printf("running completed...")
}

func (cj *cronJob) addScheduleFunc(spec string, cmd FuncJob) (cron.EntryID, error) {
	job := NewCustomJobFunc(cj, cmd)
	return cj.cron.AddJob(spec, job)
}
