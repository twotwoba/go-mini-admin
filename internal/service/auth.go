package service

import (
	"errors"
	"go-mini-admin/internal/infrastructure/jwt"
	"go-mini-admin/internal/infrastructure/response"
	"go-mini-admin/internal/model"
	"go-mini-admin/internal/repository"
	"go-mini-admin/pkg/utils"
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

var (
	ErrInvalidPassword  = errors.New("密码错误")
	ErrGenerateTokenErr = errors.New("生成令牌失败")
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token string                `json:"token"`
	User  *repository.LoginUser `json:"user"`
}

func (s *authService) Register(req *CreateUserRequest) error {
	return s.userService.CreateUser(req)
}

func (s *authService) Login(req *LoginRequest) (*LoginResponse, error) {
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("用户名或密码不能为空")
	}

	user, err := s.userRepo.GetUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New(response.GetErrMsg(response.CodeInvalidCredentials))
	}
	token, err := s.jwtManager.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}
	loginUser := &repository.LoginUser{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    *user.Email,
		Phone:    user.Phone,
		Avatar:   user.Avatar,
		Status:   user.Status,
		Roles:    user.Roles,
	}
	// TODO 是不是应该放到Redis中
	return &LoginResponse{Token: token, User: loginUser}, nil
}

func (s *authService) Logout(userID int) error {
	return nil
}

func (s *authService) GetUserInfo(userID int) (*model.User, error) {
	return s.userRepo.GetUserByID(userID)
}
