package core

import (
	"fmt"
	"go-mini-admin/config"
	"go-mini-admin/internal/core/global"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// viper 初始化并解析配置文件
func Viper() *viper.Viper {
	configFile := config.GetConfigFile()
	v := viper.New()

	if configFile != "" {
		// 通过命令行或环境变量指定了具体文件，直接加载该文件
		v.SetConfigFile(configFile)
	} else {
		// 这里根据环境变量 GMA_APP_ENV 自动探测, 默认 dev
		configName := fmt.Sprintf("config.%s", config.GetEnv())
		v.SetConfigName(configName)
		v.SetConfigType("yaml")

		v.AddConfigPath(".")
		v.AddConfigPath("..")

		// 适用于编译后的安装包
		if exePath, err := os.Executable(); err == nil {
			v.AddConfigPath(filepath.Dir(exePath))
		}
	}

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件失败: %w", err))
	}

	global.GMA_CONFIG = &config.Config{}
	if err := v.Unmarshal(global.GMA_CONFIG); err != nil {
		fmt.Printf("解析配置文件失败: %v\n", err)
	}

	setupGinMode(global.GMA_CONFIG.Server.Mode)

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件已修改:", e.Name)
		if err := v.Unmarshal(global.GMA_CONFIG); err != nil {
			fmt.Printf("配置文件变动后解析失败: %v\n", err)
		}
	})

	return v
}

func setupGinMode(mode string) {
	switch mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
