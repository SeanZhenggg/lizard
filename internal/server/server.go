package server

import (
	"github.com/gin-gonic/gin"
	"lizard/internal/app/web"
	"log"
	"time"
)

type appServer struct {
	gin     *gin.Engine `wire:"-"`
	iWebApp web.IWebApp
}

func (app *appServer) Init() {
	app.gin = gin.New()
	app.gin.Use(gin.Recovery())

	app.iWebApp.Init(app.gin)

	time.Local = time.UTC
	//_, err := time.LoadLocation()
	//if err != nil {
	//	log.Fatalf("時區設置異常：%v\n", err)
	//	return
	//}
}

func (app *appServer) Run() {
	err := app.gin.Run(":8080")
	if err != nil {
		log.Fatalf("appServer Run error：%v\n", err)
	}
}
