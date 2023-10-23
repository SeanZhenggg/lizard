package config

import (
	"fmt"
	"github.com/SeanZhenggg/go-utils/logger"
	"github.com/spf13/viper"
	"log"
	"os"
)

type IConfigEnv interface {
	GetLogConfig() logger.LogConfig
	GetDbConfig() DbConfig
}

func ProviderIConfigEnv() IConfigEnv {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		log.Fatalf("🔔🔔🔔 fatal error : APP_ENV required 🔔🔔🔔")
	}

	log.Printf("appEnv : %v, path: %s", appEnv, fmt.Sprintf("./configs/%s", appEnv))

	viper.AddConfigPath(fmt.Sprintf("./configs/%s/", appEnv))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("🔔🔔🔔 fatal error viper.ReadInConfig: %v 🔔🔔🔔", err)
	}

	var cfg configEnv
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("🔔🔔🔔 fatal error viper.Unmarshal: %v 🔔🔔🔔", err)
	}

	return &cfg
}

type configEnv struct {
	LogConfig logConfig `mapstructure:"log"`
	DbConfig  DbConfig  `mapstructure:"mongodb"`
}

type logConfig struct {
	Name  string `mapstructure:"name"`
	Env   string `mapstructure:"env"`
	Level string `mapstructure:"level"`
}

type DbConfig struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	DbName string `mapstructure:"dbName"`
}

func (c *configEnv) GetLogConfig() logger.LogConfig {
	return logger.LogConfig(c.LogConfig)
}

func (c *configEnv) GetDbConfig() DbConfig {
	return c.DbConfig
}
