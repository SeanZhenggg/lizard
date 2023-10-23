//go:build wireinject
// +build wireinject

package job

import (
	logUtils "github.com/SeanZhenggg/go-utils/logger"
	"github.com/google/wire"
	jobApp "lizard/internal/app/job"
	"lizard/internal/config"
	"lizard/internal/utils/cronjob"

	"lizard/internal/controller/job"
	"lizard/internal/mongo"
	"lizard/internal/repository"
	"lizard/internal/service"
)

func NewJobServer() *jobServer {
	panic(
		wire.Build(
			config.ProviderIConfigEnv,
			cronjob.ProviderCronJob,
			wire.Bind(new(logUtils.ILogConfig), new(config.IConfigEnv)),
			logUtils.ProviderILogger,
			repository.ProvideTrendRepository,
			service.ProviderITrendsSrv,
			job.ProviderITrendsJobCtrl,
			job.ProvideJobController,
			mongo.ProvideMongoDbCli,
			jobApp.ProvideJobApp,
			wire.Struct(new(jobServer), "*"),
		),
	)
}
