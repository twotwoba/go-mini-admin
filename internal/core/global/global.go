package global

import (
	"go-mini-admin/config"

	"github.com/spf13/viper"
)

var (
	GMA_CONFIG *config.Config // 全局配置
	GMA_VIPER  *viper.Viper   // 全局 viper 实例
)
