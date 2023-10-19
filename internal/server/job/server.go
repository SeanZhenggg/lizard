package job

import (
	"fmt"
	"lizard/internal/app/job"
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
	fmt.Printf("job run...")
}
