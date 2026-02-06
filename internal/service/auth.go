package service

import (
	"go-mini-admin/internal/infrastructure/jwt"
	"go-mini-admin/internal/model"
	"go-mini-admin/internal/repository"
)

type AuthService interface {
	Register(req *CreateUserRequest) error
	Login(req *LoginRequest) (*LoginResponse, error)
	Logout(userID int) error
	GetUserInfo(userID int) (*model.User, error)
}

type authService struct {
	userService UserService
	userRepo    repository.UserRepository
	jwtManager  *jwt.JWTManager
}

func NewAuthService(userService UserService, userRepo repository.UserRepository, jwtManager *jwt.JWTManager) AuthService {
	return &authService{
		userRepo:    userRepo,
		userService: userService,
		jwtManager:  jwtManager,
	}
}

type LoginRequest struct {
}
type LoginResponse struct {
	Token string `json:"token"`
}

func (s authService) Register(req *CreateUserRequest) error {
	return s.userService.CreateUser(req)
}

func (s authService) Login(req *LoginRequest) (*LoginResponse, error) {
	return &LoginResponse{
		Token: "sk-sadfgffas",
	}, nil
}

func (s authService) Logout(userID int) error {
	return nil
}

func (s authService) GetUserInfo(userId int) (*model.User, error) {
	return s.userRepo.GetUserByID(userId)
}
