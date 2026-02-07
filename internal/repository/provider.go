package repository

import "gorm.io/gorm"

type Provider struct {
	User UserRepository
}

func NewProvider(db *gorm.DB) *Provider {
	return &Provider{
		User: NewUserRepository(db),
	}
}
