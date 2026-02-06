package service

import (
	"errors"
	"go-mini-admin/internal/model"
	"go-mini-admin/internal/repository"
	"go-mini-admin/pkg/utils"
)

type UserService interface {
	CreateUser(req *CreateUserRequest) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=5,max=50"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Email    string `json:"email" binding:"omitempty,email"`
	Nickname string `json:"nickname" binding:"omitempty,min=5,max=50"`
	Phone    string `json:"phone" binding:"omitempty,min=11,max=11"`
	Status   int    `json:"status"`
}

func (s *userService) CreateUser(req *CreateUserRequest) error {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Status:   1,
	}
	if err := s.userRepo.CreateUser(user); err != nil {
		if utils.IsDuplicateError(err) {
			return errors.New("用户名或邮箱已存在")
		}
		return err
	}
	return nil
}
