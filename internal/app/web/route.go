package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *webApp) setApiRoutes(g *gin.Engine) {
	g.GET("health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "ok"}) })

	webGroup := g.Group("")
	webGroup.Use(app.RespMw.Handle)
	webGroup.GET("/r/:path", app.Ctrl.TrendCtrl.RedirectToTrendPage)
	webGroup.POST("/line/webhook", app.Ctrl.MessageCtrl.RecvMessage)

	apiGroup := g.Group("/api")
	apiGroup.Use(app.RespMw.Handle)
	apiGroup.POST("trend/push-msg", app.Ctrl.TrendCtrl.FetchTrendsAndPushMessage)

}
