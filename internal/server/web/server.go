package web

import (
	"github.com/gin-gonic/gin"
	"lizard/internal/app/web"
	"log"
)

type appServer struct {
	gin     *gin.Engine `wire:"-"`
	iWebApp web.IWebApp
}

func (app *appServer) Init() {
	app.gin = gin.New()
	app.gin.Use(gin.Recovery())

	app.iWebApp.Init(app.gin)
}

func (app *appServer) Run() {
	err := app.gin.Run(":8080")
	if err != nil {
		log.Fatalf("appServer Run error：%v\n", err)
	}
}
