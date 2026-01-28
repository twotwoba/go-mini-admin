package model

// permissions表
type Permission struct {
	BaseModel
	Name        string `json:"name" gorm:"size:50;not null"`
	Code        string `json:"code" gorm:"size:100;uniqueIndex;not null;comment:permission code e.g. user:create"`
	Type        int    `json:"type" gorm:"default:1;comment:1-menu 2-button 3-api"`
	ParentID    uint   `json:"parent_id" gorm:"default:0"`
	Path        string `json:"path" gorm:"size:255;comment:API path"`
	Method      string `json:"method" gorm:"size:20;comment:HTTP method"`
	Icon        string `json:"icon" gorm:"size:50"`
	Sort        int    `json:"sort" gorm:"default:0"`
	Status      int    `json:"status" gorm:"default:1;comment:1-enabled 0-disabled"`
	Description string `json:"description" gorm:"size:255"`
}

func (Permission) TableName() string {
	return "permissions"
}

// 权限类型
const (
	PermissionTypeMenu   = 1
	PermissionTypeButton = 2
	PermissionTypeAPI    = 3
)
