package web

import (
	"lizard/internal/server/web"
)

func main() {

	appServer := web.NewAppServer()
	appServer.Init()
	appServer.Run()
}
