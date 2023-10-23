package web

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"net/http"
)

func (app *webApp) setApiRoutes(g *gin.Engine) {
	group := g.Group("/api")
	group.Use(app.RespMw.ResponseHandler)

	group.GET("health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "ok"}) })
	group.GET("fetchTrends", app.Ctrl.TrendCtrl.FetchTrends)

	g.POST("/line/webhook", func(ctx *gin.Context) {
		bot, err := linebot.New("d48ec0332fb59de64035f2c555bb9995", "aAV/cg5zZd3nY19Po6aawPs7C1vBotu+AJHs0wMdndKbrm5XPqQ4XOYUa1QSYrDBfvLK+t0JU8v1cf1BQbZvepXinztFKns2xa79JZEbTbRbRPUqcuZXbYY5FF5eAEyJmAR8msweC0yumXr11Pz62QdB04t89/1O/w1cDnyilFU=")
		if err != nil {
			log.Printf("linebot.New error : %v\n", err)
			return
		}

		events, err := bot.ParseRequest(ctx.Request)
		if err != nil {
			log.Printf("bot.ParseRequest error : %v\n", err)
			return
		}

		for _, event := range events {
			log.Printf("event source : %+v\n", event.Source)
		}
		ctx.Status(http.StatusOK)
	})
}
