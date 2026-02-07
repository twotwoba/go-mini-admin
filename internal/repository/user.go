package repository

import (
	"errors"
	"fmt"
	"go-mini-admin/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
	ExistsByUsername(username string) (bool, error)
	ExistsByEmail(e string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

var (
	ErrUserNotFound = errors.New("用户不存在")
	ErrUserDisabled = errors.New("用户已被禁用")
)

func (r *userRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id int) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *userRepository) GetUserByID(id int) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

type LoginUser struct {
	ID       uint         `json:"id"`
	Username string       `json:"username"`
	Nickname string       `json:"nickname"`
	Email    string       `json:"email"`
	Phone    string       `json:"phone"`
	Avatar   string       `json:"avatar"`
	Status   int          `json:"status"`
	Roles    []model.Role `json:"roles"`
}

func (r *userRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.db.Preload("Roles").Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound // 用户不存在返回明确错误
		}
		return nil, fmt.Errorf("数据库查询失败：%w", err)
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) ExistsByUsername(username string) (bool, error) {
	var count int64
	if err := r.db.Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
func (r *userRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	if err := r.db.Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
