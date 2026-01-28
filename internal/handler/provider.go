package handler

import "go-mini-admin/internal/service"

type Handlers struct {
}

func NewHandlers(service *service.Services) *Handlers {
	return &Handlers{}
}
