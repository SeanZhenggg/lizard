package job

import (
	"lizard/internal/app/job"
	"log"
)

type jobServer struct {
	iJobApp job.IJobApp
}

func (job *jobServer) Init() {
	// TODO
	job.iJobApp.Init()

}

func (job *jobServer) Run() {
	// TODO
	log.Printf("jobServer Start")
	job.iJobApp.Start()
}
