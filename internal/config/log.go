package config

import (
	"github.com/SeanZhenggg/go-utils/logger"
)

func ProviderILogConfig(env IConfigEnv) logger.ILogConfig {
	return &config{
		Name:  env.GetLogConfig().Name,
		Env:   env.GetLogConfig().Env,
		Level: env.GetLogConfig().Level,
	}
}

type config struct {
	Name  string
	Env   string
	Level string
}

func (c *config) GetLogConfig() logger.LogConfig {
	return logger.LogConfig(*c)
}
