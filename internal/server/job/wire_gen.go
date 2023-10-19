// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package job

import (
	"github.com/SeanZhenggg/go-utils/logger"
	job2 "lizard/internal/app/job"
	"lizard/internal/config"
	"lizard/internal/controller/job"
	"lizard/internal/mongo"
	"lizard/internal/repository"
	"lizard/internal/service"
)

// Injectors from wire.go:

func NewJobServer() *jobServer {
	iConfigEnv := config.ProviderIConfigEnv()
	iLogConfig := config.ProviderILogConfig(iConfigEnv)
	iLogger := logger.ProviderILogger(iLogConfig)
	iMongoCli := mongo.ProvideMongoDbCli(iConfigEnv)
	iTrendRepository := repository.ProvideTrendRepository()
	iTrendSrv := service.ProviderITrendsSrv(iLogger, iMongoCli, iTrendRepository)
	iTrendJobCtrl := job.ProviderITrendsJobCtrl(iTrendSrv)
	controller := job.ProvideJobController(iTrendJobCtrl)
	iJobApp := job2.ProvideJobApp(controller)
	webAppServer := &jobServer{
		iJobApp: iJobApp,
	}
	return webAppServer
}