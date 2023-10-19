//go:build wireinject
// +build wireinject

package server

import (
	logUtils "github.com/SeanZhenggg/go-utils/logger"
	"github.com/google/wire"
	"lizard/source/app/web"
	"lizard/source/config"
	web2 "lizard/source/controllers/web"
	middleware2 "lizard/source/controllers/web/middleware"
	"lizard/source/mongo"
	"lizard/source/repository"
	"lizard/source/service"
)

func NewAppServer() *appServer {
	panic(
		wire.Build(
			config.ProviderIConfigEnv,
			config.ProviderILogConfig,
			logUtils.ProviderILogger,
			repository.ProvideTrendRepository,
			middleware2.ProvideResponseMiddleware,
			middleware2.ProvideAuthMiddleware,
			service.ProviderITrendsSrv,
			web2.ProviderITrendsCtrl,
			web2.ProvideController,
			mongo.ProvideMongoDbCli,
			web.ProvideWebApp,
			wire.Struct(new(appServer), "*"),
		),
	)
}
