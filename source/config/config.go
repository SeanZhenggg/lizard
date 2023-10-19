package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type IConfigEnv interface {
	GetLogConfig() logConfig
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

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("🔔🔔🔔 fatal error viper.ReadInConfig: %v 🔔🔔🔔", err)
	}
	log.Printf("%v", viper.Get("log"))
	// 将配置映射到结构体
	var cfg configEnv
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("🔔🔔🔔 fatal error viper.Unmarshal: %v 🔔🔔🔔", err)
	}
	log.Printf("cfg: %v", cfg)

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

func (c *configEnv) GetLogConfig() logConfig {
	return c.LogConfig
}

func (c *configEnv) GetDbConfig() DbConfig {
	return c.DbConfig
}
