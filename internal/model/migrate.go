package model

import "gorm.io/gorm"

func AllModels() []any {
	return []any{
		&User{},
		&Role{},
		&Permission{},
		&UserRole{},
		&RolePermission{},
	}
}

func AutoMigrate(db *gorm.DB, models ...any) error {
	return db.AutoMigrate(models...)
}
