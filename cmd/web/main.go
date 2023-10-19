package web

import "lizard/internal/server"

func main() {

	appServer := server.NewAppServer()
	appServer.Init()
	appServer.Run()
}
