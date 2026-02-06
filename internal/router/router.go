package router

import (
	"go-mini-admin/internal/handler"
	"go-mini-admin/internal/infrastructure/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type IRouter interface {
	RegisterRoutes(rg *gin.RouterGroup)
}

func NewRoutes(handlers *handler.Handlers, mw *middleware.Middleware) []IRouter {
	return []IRouter{
		&authRouter{handlers.Auth, mw},
	}
}

func Setup(mode string, handlers *handler.Handlers, mw *middleware.Middleware) *gin.Engine {
	gin.SetMode(mode)
	r := gin.New()

	r.Use(mw.Logger())
	r.Use(mw.Recovery())

	corsConfig := cors.DefaultConfig()
	if mode == gin.DebugMode {
		corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	}
	r.Use(cors.New(corsConfig))

	v1 := r.Group("/api/v1")

	modules := NewRoutes(handlers, mw)

	for _, module := range modules {
		module.RegisterRoutes(v1)
	}

	return r
}
