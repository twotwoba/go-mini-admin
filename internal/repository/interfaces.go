package repository

import "go-mini-admin/internal/model"

type UserRepository interface {
	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
}
