package web

import (
	"github.com/gin-gonic/gin"
	"lizard/internal/app/web"
	"lizard/internal/config"
	"log"
	"time"
)

type appServer struct {
	gin       *gin.Engine `wire:"-"`
	iWebApp   web.IWebApp
	ConfigEnv config.IConfigEnv
}

func (app *appServer) Init() {
	app.gin = gin.New()
	app.gin.Use(gin.Recovery())
	time.Local = time.UTC
	app.iWebApp.Init(app.gin)
}

func (app *appServer) Run() {
	port := app.ConfigEnv.GetHttpConfig().Port
	address := ":" + port
	err := app.gin.Run(address)
	if err != nil {
		log.Fatalf("appServer Run error：%v\n", err)
	}
}
