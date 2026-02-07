package service

import (
	"go-mini-admin/internal/infrastructure/jwt"
	"go-mini-admin/internal/repository"
)

type Provider struct {
	Auth AuthService
	User UserService
}

func NewProvider(repository *repository.Provider, jwtManager *jwt.JWTManager) *Provider {
	userService := NewUserService(repository.User)
	authService := NewAuthService(userService, repository.User, jwtManager)
	return &Provider{
		Auth: authService,
		User: userService,
	}
}
