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
			// update : 這裡告訴 wire IConfigEnv 可以當作 ILogConfig
			// 下面就能正確注入 ILogConfig 到 logUtils.ProviderILogger
			wire.Bind(new(logUtils.ILogConfig), new(config.IConfigEnv)),
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
