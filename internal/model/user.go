package model

// users 表
type User struct {
	BaseModel
	Username string  `json:"username" gorm:"size:50;uniqueIndex;not null;comment:用户名"`
	Password string  `json:"-" gorm:"size:255;not null"`
	Nickname string  `json:"nickname" gorm:"size:50"`
	Email    *string `json:"email" gorm:"size:100;uniqueIndex"`
	Phone    string  `json:"phone" gorm:"size:20";defautl:'';comment:"手机号"`
	Avatar   string  `json:"avatar" gorm:"size:255"`
	Status   int     `json:"status" gorm:"default:1;comment:1-enabled 0-disabled"`
	Roles    []Role  `json:"roles" gorm:"many2many:user_roles;"`
}

func (User) TableName() string {
	return "users"
}

// user_roles 表
type UserRole struct {
	UserID uint `gorm:"primaryKey"`
	RoleID uint `gorm:"primaryKey"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
