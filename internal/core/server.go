package core

import (
	"go-mini-admin/internal/core/global"
)

func Run() {
	initSystem()
	// ServerRun()
}

// 系统初始化
func initSystem() {
	global.GMA_VIPER = Viper()
}
