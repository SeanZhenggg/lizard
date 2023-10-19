package web

import (
	"lizard/internal/app/job"
)

type appServer struct {
	iJobApp job.IJobApp
}

func (app *appServer) Init() {
}

func (app *appServer) Run() {
}
