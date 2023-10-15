package config

import (
	"fmt"
	"github.com/SeanZhenggg/go-utils/logger"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type IConfigEnv interface {
	logger.ILogConfig
}

func ProviderIConfigEnv() logger.ILogConfig {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// 获取可执行文件所在目录
	exPath := filepath.Dir(ex)

	// 获取项目根目录（假设根目录为 src 的上一级）
	projectRoot := filepath.Dir(exPath)

	// 打印项目根目录
	//println("Project Root:", projectRoot)

	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(projectRoot)

	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 将配置映射到结构体
	var config ConfigEnv
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct：%s", err))
	}

	return &config
}

type ConfigEnv struct {
	LogConfig LogConfig
}

type LogConfig struct {
	Name  string `mapstructure:"name"`
	Env   string `mapstructure:"env"`
	Level string `mapstructure:"level"`
}

func (c ConfigEnv) GetLogConfig() logger.LogConfig {
	return logger.LogConfig{
		Name:  c.LogConfig.Name,
		Env:   c.LogConfig.Env,
		Level: c.LogConfig.Level,
	}
}
