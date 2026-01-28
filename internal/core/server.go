package core

import (
	"fmt"
	"go-mini-admin/config"
	"os"
)

// 系统入口
func Run() {
	initSystem()
	// ServerRun()
}

// 系统初始化
func initSystem() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("加载配置文件失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("cfg: %v\n", cfg)
}
