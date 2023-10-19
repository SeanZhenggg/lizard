package main

import "lizard/internal/server/job"

func main() {
	jobApp := job.NewJobServer()
	jobApp.Init()
	jobApp.Run()
}
