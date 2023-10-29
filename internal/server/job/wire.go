//go:build wireinject
// +build wireinject

package job

import (
	logUtils "github.com/SeanZhenggg/go-utils/logger"
	"github.com/google/wire"
	jobApp "lizard/internal/app/job"
	"lizard/internal/config"
	"lizard/internal/utils/cronjob"

	jobCtrl "lizard/internal/controller/job"
	jobMw "lizard/internal/controller/job/middleware"
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
			jobMw.ProvideJobLogMiddleware,
			repository.ProvideTrendRepository,
			service.ProviderITrendsSrv,
			service.ProvideMessageSrv,
			jobCtrl.ProviderITrendsJobCtrl,
			jobCtrl.ProvideJobController,
			mongo.ProvideMongoDbCli,
			jobApp.ProvideJobApp,
			wire.Struct(new(jobServer), "*"),
		),
	)
}
