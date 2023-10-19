//go:build wireinject
// +build wireinject

package server

import (
	logUtils "github.com/SeanZhenggg/go-utils/logger"
	"github.com/google/wire"
	"lizard/internal/app/web"
	"lizard/internal/config"
	web2 "lizard/internal/controller/web"
	middleware2 "lizard/internal/controller/web/middleware"
	"lizard/internal/mongo"
	"lizard/internal/repository"
	"lizard/internal/service"
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
