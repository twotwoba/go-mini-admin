package core

import (
	"fmt"
	"go-mini-admin/config"
	"go-mini-admin/internal/handler"
	"go-mini-admin/internal/infrastructure/logger"
	"go-mini-admin/internal/infrastructure/middleware"
	"go-mini-admin/internal/router"
)

func ServerRun(cfg *config.Config, handlers *handler.Provider, mw *middleware.Middleware) {
	r := router.Setup(cfg.Server.Mode, handlers, mw)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := r.Run(addr); err != nil {
		logger.Fatalf("❌服务启动失败: %v", err)
	}
}
