package handler

import (
	"go-mini-admin/internal/infrastructure/response"
	"go-mini-admin/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	// TODO
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req service.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, err.Error())
		return
	}

	if err := h.authService.Register(&req); err != nil {
		response.Fail(c, err.Error())
		return
	}
	response.Success(c, "用户注册成功")
}
