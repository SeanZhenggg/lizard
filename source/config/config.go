package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type IConfigEnv interface {
	GetLogConfig() logConfig
}

func ProviderIConfigEnv() IConfigEnv {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("🔔🔔🔔 fatal error os.Executable: %v 🔔🔔🔔", err)
	}

	// 获取可执行文件所在目录的路径
	exeDir := filepath.Dir(exePath)

	// 获取项目根目录的路径（假设项目根目录就在可执行文件所在的目录下）
	projectRoot := filepath.Dir(exeDir)

	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(projectRoot)

	// 读取配置文件
	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("🔔🔔🔔 fatal error viper.ReadInConfig: %v 🔔🔔🔔", err)
	}

	// 将配置映射到结构体
	var config configEnv
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("🔔🔔🔔 fatal error viper.Unmarshal: %v 🔔🔔🔔", err)
	}

	return &config
}

type configEnv struct {
	LogConfig logConfig
}

type logConfig struct {
	Name  string `mapstructure:"name"`
	Env   string `mapstructure:"env"`
	Level string `mapstructure:"level"`
}

func (c *configEnv) GetLogConfig() logConfig {
	return c.LogConfig
}
