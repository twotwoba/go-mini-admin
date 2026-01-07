package config

import (
	"flag"
	"fmt"
	"os"
)

// 配置文件格式入口
type Config struct {
	Server Server `mapstructure:"server"`
	Zap    Zap    `mapstructure:"zap"`
}

// GetConfigFile 获取配置文件
// 优先级：命令行参数 > 环境变量 > 默认空 (由调用方处理默认逻辑)
func GetConfigFile() (configFile string) {
	flag.StringVar(&configFile, "c", "", "choose config file.")
	flag.Parse()
	if configFile != "" {
		fmt.Printf("您正在使用命令行 '-c' 参数指定配置文件: %s\n", configFile)
		return
	}

	if configFile = GetConfigPath(); configFile != "" {
		fmt.Printf("您正在使用环境变量 '%s' 指定配置文件: %s\n", GMA_APP_ENV_FILE, configFile)
		return
	}

	return
}

func GetConfigPath() string {
	return os.Getenv(GMA_APP_ENV_FILE)
}
