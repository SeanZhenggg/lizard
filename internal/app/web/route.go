package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *webApp) setApiRoutes(g *gin.Engine) {
	group := g.Group("/api")
	group.Use(app.RespMw.ResponseHandler)

	group.GET("health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "ok"}) })
	group.GET("fetchTrends", app.Ctrl.TrendCtrl.FetchTrends)

	g.POST("/line/webhook", app.Ctrl.MessageCtrl.RecvMessage)
}
