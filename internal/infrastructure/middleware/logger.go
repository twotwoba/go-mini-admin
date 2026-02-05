package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
)

type responseWrite struct {
	gin.ResponseWriter               // 包装原始 gin.ResponseWriter
	body               *bytes.Buffer // 用于缓存响应体的缓冲区
}

// 重写 gin.ResponseWriter 的 Write 方法
func (rw responseWrite) Write(data []byte) (int, error) {
	rw.body.Write(data)
	return rw.ResponseWriter.Write(data)
}

func (mw *Middleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		/**
		 * 其他业务处理
		 * ......
		 */

		rw := &responseWrite{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = rw

		c.Next()

		latency := time.Since(start)
		mw.logger.Infof("| %3d | %13v | %15s | %-7s %s%s",
			c.Writer.Status(),
			latency,
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.URL.RawQuery)
	}
}
