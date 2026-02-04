package middleware

import (
	"go-mini-admin/internal/infrastructure/jwt"
	"go-mini-admin/internal/infrastructure/logger"
	"go-mini-admin/internal/repository"
)

// 中间件的依赖
type Middleware struct {
	logger     logger.Logger
	userRepo   repository.UserRepository
	JWTManager *jwt.JWTManager
}

func New(logger logger.Logger, userRepo repository.UserRepository, jwtManager *jwt.JWTManager) *Middleware {
	return &Middleware{
		logger:     logger,
		userRepo:   userRepo,
		JWTManager: jwtManager,
	}
}
