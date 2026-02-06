package service

import (
	"go-mini-admin/internal/infrastructure/jwt"
	"go-mini-admin/internal/repository"
)

type Services struct {
	Auth AuthService
	User UserService
}

func NewServices(repository *repository.Repositories, jwtManager *jwt.JWTManager) *Services {
	userService := NewUserService(repository.User)
	return &Services{
		Auth: NewAuthService(userService, repository.User, jwtManager),
		User: userService,
	}
}
