package core

import (
	"context"
	"errors"
	"fmt"
	"go-mini-admin/config"
	"go-mini-admin/internal/handler"
	"go-mini-admin/internal/infrastructure/logger"
	"go-mini-admin/internal/infrastructure/middleware"
	"go-mini-admin/internal/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ServerRun(cfg *config.Config, handlers *handler.Provider, mw *middleware.Middleware) {
	r := router.Setup(cfg.Server.Mode, handlers, mw)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: r,
	}

	// 在 goroutine 中启动服务
	go func() {
		logger.Infof("\n🚀 服务启动，监听 %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("\n❌ 服务启动失败: %v", err)
		}
	}()

	// 等待中断信号（SIGINT / SIGTERM）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	logger.Infof("⏳ 收到信号 %v，开始优雅关闭...", sig)

	// 给予 10 秒超时，等待已有请求处理完毕
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("❌ 服务关闭异常: %v", err)
	} else {
		logger.Info("✅ 服务已优雅关闭")
	}
}
