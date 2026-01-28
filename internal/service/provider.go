package service

import (
	"go-mini-admin/internal/infrastructure/jwt"
	"go-mini-admin/internal/repository"
)

type Services struct {
}

func NewServices(repository *repository.Repositories, jwtManager *jwt.JWTManager) *Services {
	return &Services{}
}
