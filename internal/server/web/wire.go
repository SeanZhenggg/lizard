//go:build wireinject
// +build wireinject

package web

import (
	logUtils "github.com/SeanZhenggg/go-utils/logger"
	"github.com/google/wire"
	"lizard/internal/app/web"
	"lizard/internal/config"
	webApp "lizard/internal/controller/web"
	"lizard/internal/controller/web/middleware"
	"lizard/internal/mongo"
	"lizard/internal/repository"
	"lizard/internal/service"
)

func NewAppServer() *appServer {
	panic(
		wire.Build(
			config.ProviderIConfigEnv,
			wire.Bind(new(logUtils.ILogConfig), new(config.IConfigEnv)),
			logUtils.ProviderILogger,
			repository.ProvideTrendRepository,
			middleware.ProvideResponseMiddleware,
			middleware.ProvideAuthMiddleware,
			service.ProviderITrendsSrv,
			webApp.ProviderITrendsCtrl,
			webApp.ProvideMessageCtrl,
			webApp.ProvideController,
			mongo.ProvideMongoDbCli,
			web.ProvideWebApp,
			wire.Struct(new(appServer), "*"),
		),
	)
}
