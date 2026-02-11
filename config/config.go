package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config 配置文件格式入口
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Jwt      JwtConfig      `mapstructure:"jwt"`
	Log      LogConfig      `mapstructure:"log"`
}

func Load() (*Config, error) {
	configFile := getConfigFile()
	if configFile == "" {
		// 开发环境自动探测配置文件
		configFile = findDevConfig()
		if configFile == "" {
			return nil, fmt.Errorf("未找到配置文件，请通过 -c 参数或 %s 环境变量指定", GMA_APP_ENV_FILE)
		}
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

// getConfigFile 获取配置文件
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

// findDevConfig 在常见位置查找开发环境配置文件
func findDevConfig() string {
	candidates := []string{
		"config.dev.yaml",    // 当前目录 (GoLand 默认在项目根目录执行)
		"./config.dev.yaml",  // 当前目录
		"../config.dev.yaml", // 上级目录 (从 cmd/ 子目录执行时)
	}

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return ""
}
