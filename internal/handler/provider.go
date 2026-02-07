package handler

import "go-mini-admin/internal/service"

type Provider struct {
	Auth *AuthHandler
}

func NewProvider(service *service.Provider) *Provider {
	return &Provider{
		Auth: NewAuthHandler(service.Auth),
	}
}
