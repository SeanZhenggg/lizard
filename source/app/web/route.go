package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *webApp) setApiRoutes(g *gin.Engine) {
	group := g.Group("/api")

	group.GET("test", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "Hello World!"}) })
	group.GET("trends", app.Ctrl.TrendCtrl.GetTrends)
}
