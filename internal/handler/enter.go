package handler

import "go-mini-admin/internal/service"

type Handlers struct {
	Auth *AuthHandler
}

func NewHandlers(service *service.Services) *Handlers {
	return &Handlers{
		Auth: NewAuthHandler(service.Auth),
	}
}
