package model

// roles 表
type Role struct {
	BaseModel
	Name        string       `json:"name" gorm:"size:50;uniqueIndex;not null"`
	Code        string       `json:"code" gorm:"size:50;uniqueIndex;not null"`
	Description string       `json:"description" gorm:"size:255"`
	Status      int          `json:"status" gorm:"default:1;comment:1-enabled 0-disabled"`
	Sort        int          `json:"sort" gorm:"default:0"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"`
	Users       []User       `json:"users,omitempty" gorm:"many2many:user_roles;"`
}

func (Role) TableName() string {
	return "roles"
}

// role_permissions 表
type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
