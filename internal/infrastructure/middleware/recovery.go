package middleware

import (
	"go-mini-admin/internal/infrastructure/response"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (m *Middleware) Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				m.logger.Error(
					"panic recovered",                      // 日志主题
					zap.Any("error", err),                  // 结构化错误字段
					zap.ByteString("stack", debug.Stack()), // 调用栈字段（[]byte 类型适配）
				)
				response.Result(c, http.StatusInternalServerError, response.CodeServerError, "internal server error", nil)
				c.Abort()
			}
		}()
		c.Next()
	}
}
