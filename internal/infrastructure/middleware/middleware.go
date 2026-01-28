package middleware

import (
	"go-mini-admin/internal/infrastructure/jwt"
	"go-mini-admin/internal/infrastructure/logger"
	"go-mini-admin/internal/repository"
)

type Middleware struct {
}

func New(logger logger.Logger, userRepo repository.UserRepository, jwtManager *jwt.JWTManager) *Middleware {
	return &Middleware{}
}
