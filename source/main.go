package main

import "lizard/source/server"

func main() {

	appServer := server.NewAppServer()
	appServer.Init()
	appServer.Run()
}
