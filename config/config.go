package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// 配置文件格式入口
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Jwt      JwtConfig      `mapstructure:"jwt"`
	Log      LogConfig      `mapstructure:"log"`
}

func Load() (*Config, error) {
	configFile := getConfigFile()
	if configFile == "" {
		// 默认使用 dev 环境，
		configFile = "../config.dev.yaml"
		fmt.Printf("您正在使用默认配置文件: %s\n", configFile)
	}

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}
	return &cfg, nil
}

// GetConfigFile 获取配置文件
// 优先级：命令行参数 > 环境变量 > 默认空 (由调用方处理默认逻辑)
func getConfigFile() (configFile string) {
	flag.StringVar(&configFile, "c", "", "choose config file.")
	flag.Parse()
	if configFile != "" {
		fmt.Printf("您正在使用命令行 '-c' 参数指定配置文件: %s\n", configFile)
		return
	}

	if configFile = os.Getenv(GMA_APP_ENV_FILE); configFile != "" {
		fmt.Printf("您正在使用环境变量 '%s' 指定配置文件: %s\n", GMA_APP_ENV_FILE, configFile)
		return
	}

	return
}
