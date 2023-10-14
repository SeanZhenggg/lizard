//go:build wireinject
// +build wireinject

package server

import (
	logUtils "github.com/SeanZhenggg/go-utils/logger"
	"github.com/google/wire"
	"lizard/source/app/web"
	"lizard/source/controllers"
	"lizard/source/controllers/middleware"
	"lizard/source/mongo"
)

func NewAppServer() *appServer {
	panic(
		wire.Build(
			logUtils.ProviderLogger,
			middleware.ProvideResponseMiddleware,
			middleware.ProvideAuthMiddleware,
			controllers.ProviderITrendsCtrl,
			controllers.ProvideController,
			mongo.ProvideMongoDbCli,
			web.ProvideWebApp,
			wire.Struct(new(appServer), "*"),
		),
	)
}
