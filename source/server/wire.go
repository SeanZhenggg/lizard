//go:build wireinject
// +build wireinject

package server

import (
	logUtils "github.com/SeanZhenggg/go-utils/logger"
	"github.com/google/wire"
	"lizard/source/app/web"
	"lizard/source/config"
	"lizard/source/controllers"
	"lizard/source/controllers/middleware"
	"lizard/source/mongo"
	"lizard/source/repository"
	"lizard/source/service"
)

func NewAppServer() *appServer {
	panic(
		wire.Build(
			config.ProviderIConfigEnv,
			// 沒有這個的話 ProviderILogger 在 wire gen 會報錯
			//config.ProviderILogConfig,
			logUtils.ProviderILogger,
			repository.ProvideTrendRepository,
			middleware.ProvideResponseMiddleware,
			middleware.ProvideAuthMiddleware,
			service.ProviderITrendsSrv,
			controllers.ProviderITrendsCtrl,
			controllers.ProvideController,
			mongo.ProvideMongoDbCli,
			web.ProvideWebApp,
			wire.Struct(new(appServer), "*"),
		),
	)
}
