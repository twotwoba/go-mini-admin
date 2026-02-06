package router

import (
	"go-mini-admin/internal/handler"
	"go-mini-admin/internal/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

type authRouter struct {
	handlers *handler.AuthHandler
	mw       *middleware.Middleware
}

func (r authRouter) RegisterRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", r.handlers.Login)
		auth.POST("/register", r.handlers.Register)
	}
}
