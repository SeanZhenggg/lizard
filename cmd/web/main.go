package main

import (
	"lizard/internal/server/job"
	"lizard/internal/server/web"
)

func main() {
	jobApp := job.NewJobServer()
	jobApp.Init()
	jobApp.Run()

	appServer := web.NewAppServer()
	appServer.Init()
	appServer.Run()
}
