package job

import (
	"lizard/internal/app/job"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type jobServer struct {
	iJobApp job.IJobApp
}

func (job *jobServer) Init() {
	job.iJobApp.Init()
}

func (job *jobServer) Run() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	job.iJobApp.Start()

	select {
	case <-interrupt:
		log.Printf("gracefully shutdown the job service...")
		ctx := job.iJobApp.Stop()
		<-ctx.Done()
		log.Printf("service shutted down.")
	}
}
