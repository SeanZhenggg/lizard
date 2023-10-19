//go:build wireinject
// +build wireinject

package web

import (
	logUtils "github.com/SeanZhenggg/go-utils/logger"
	"github.com/google/wire"
	jobApp "lizard/internal/app/job"
	"lizard/internal/config"

	"lizard/internal/controller/job"
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
			service.ProviderITrendsSrv,
			job.ProviderITrendsJobCtrl,
			job.ProvideJobController,
			mongo.ProvideMongoDbCli,
			jobApp.ProvideJobApp,
			wire.Struct(new(appServer), "*"),
		),
	)
}
