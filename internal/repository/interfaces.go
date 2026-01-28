package repository

import (
	"go-mini-admin/internal/handler"
	"go-mini-admin/internal/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(mode string, handlers *handler.Handlers, mw *middleware.Middleware) *gin.Engine {
	r := gin.Default()
	return r
}
