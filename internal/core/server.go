package core

import (
	"fmt"
	"go-mini-admin/config"
	"go-mini-admin/internal/handler"
	"go-mini-admin/internal/infrastructure/database"
	"go-mini-admin/internal/infrastructure/jwt"
	"go-mini-admin/internal/infrastructure/logger"
	"go-mini-admin/internal/infrastructure/middleware"
	"go-mini-admin/internal/model"
	"go-mini-admin/internal/repository"
	"go-mini-admin/internal/service"
	"go-mini-admin/pkg/utils"
	"os"

	"gorm.io/gorm"
)

// 系统入口
func Run() {
	initSystem()
}

// 系统初始化
func initSystem() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("加载配置文件失败: %v\n", err)
		os.Exit(1)
	}
	// fmt.Printf("cfg: %+v\n", cfg)

	if err := logger.Init(&cfg.Log); err != nil {
		fmt.Printf("⭕初始化日志失败：%v\n", err)
		os.Exit(1)
	}
	logger.Info("✅日志初始化完成")

	db, err := database.NewMysql(&cfg.Database, cfg.Server.Mode)
	if err != nil {
		fmt.Printf("⭕初始化数据库失败：%v\n", err)
		os.Exit(1)
	}
	logger.Info("✅数据库已连接")
	if err := model.AutoMigrate(db, model.AllModels()...); err != nil {
		logger.Fatalf("⭕数据库自动迁移失败：%v", err)
	}
	logger.Info("✅数据库迁移完成")

	jwtManager := jwt.New(&cfg.Jwt)

	repos := repository.NewProvider(db)
	services := service.NewProvider(repos, jwtManager)
	handlers := handler.NewProvider(services)
	mw := middleware.New(logger.Default(), repos.User, jwtManager)

	initDefaultData(db)

	ServerRun(cfg, handlers, mw)

	// 服务关闭后，清理资源
	database.Close(db)
	logger.Info("✅ 数据库连接已关闭")
	logger.Info("👋 服务已退出")
}

func initDefaultData(db *gorm.DB) {
	// 查看 admin 角色是否存在
	var count int64
	db.Model(&model.Role{}).Where("code = ?", "admin").Count(&count)
	if count > 0 {
		return
	}

	logger.Info("初始化默认数据...")

	// 创建默认权限
	permissions := []model.Permission{
		{Name: "用户管理", Code: "user", Type: 1, ParentID: 0, Sort: 1, Status: 1},
		{Name: "用户列表", Code: "user:list", Type: 3, ParentID: 1, Path: "/api/v1/users", Method: "GET", Sort: 1, Status: 1},
		{Name: "创建用户", Code: "user:create", Type: 3, ParentID: 1, Path: "/api/v1/users", Method: "POST", Sort: 2, Status: 1},
		{Name: "查看用户", Code: "user:read", Type: 3, ParentID: 1, Path: "/api/v1/users/:id", Method: "GET", Sort: 3, Status: 1},
		{Name: "更新用户", Code: "user:update", Type: 3, ParentID: 1, Path: "/api/v1/users/:id", Method: "PUT", Sort: 4, Status: 1},
		{Name: "删除用户", Code: "user:delete", Type: 3, ParentID: 1, Path: "/api/v1/users/:id", Method: "DELETE", Sort: 5, Status: 1},
		{Name: "分配角色", Code: "user:assign-roles", Type: 3, ParentID: 1, Path: "/api/v1/users/:id/roles", Method: "PUT", Sort: 6, Status: 1},
		{Name: "角色管理", Code: "role", Type: 1, ParentID: 0, Sort: 2, Status: 1},
		{Name: "角色列表", Code: "role:list", Type: 3, ParentID: 8, Path: "/api/v1/roles", Method: "GET", Sort: 1, Status: 1},
		{Name: "创建角色", Code: "role:create", Type: 3, ParentID: 8, Path: "/api/v1/roles", Method: "POST", Sort: 2, Status: 1},
		{Name: "查看角色", Code: "role:read", Type: 3, ParentID: 8, Path: "/api/v1/roles/:id", Method: "GET", Sort: 3, Status: 1},
		{Name: "更新角色", Code: "role:update", Type: 3, ParentID: 8, Path: "/api/v1/roles/:id", Method: "PUT", Sort: 4, Status: 1},
		{Name: "删除角色", Code: "role:delete", Type: 3, ParentID: 8, Path: "/api/v1/roles/:id", Method: "DELETE", Sort: 5, Status: 1},
		{Name: "分配权限", Code: "role:assign-permissions", Type: 3, ParentID: 8, Path: "/api/v1/roles/:id/permissions", Method: "PUT", Sort: 6, Status: 1},
		{Name: "权限管理", Code: "permission", Type: 1, ParentID: 0, Sort: 3, Status: 1},
		{Name: "权限列表", Code: "permission:list", Type: 3, ParentID: 15, Path: "/api/v1/permissions", Method: "GET", Sort: 1, Status: 1},
		{Name: "创建权限", Code: "permission:create", Type: 3, ParentID: 15, Path: "/api/v1/permissions", Method: "POST", Sort: 2, Status: 1},
		{Name: "查看权限", Code: "permission:read", Type: 3, ParentID: 15, Path: "/api/v1/permissions/:id", Method: "GET", Sort: 3, Status: 1},
		{Name: "更新权限", Code: "permission:update", Type: 3, ParentID: 15, Path: "/api/v1/permissions/:id", Method: "PUT", Sort: 4, Status: 1},
		{Name: "删除权限", Code: "permission:delete", Type: 3, ParentID: 15, Path: "/api/v1/permissions/:id", Method: "DELETE", Sort: 5, Status: 1},
	}

	for i := range permissions {
		db.Create(&permissions[i])
	}

	// 创建 admin role
	adminRole := model.Role{
		Name:        "超级管理员",
		Code:        "admin",
		Description: "系统超级管理员，拥有所有权限",
		Status:      1,
		Sort:        0,
	}
	db.Create(&adminRole)

	// 创建 admin user
	adminUser := model.User{
		Username: "admin",
		Password: "$2a$10$N9qo8uLOickgx2ZMRZoMy.MqrqPiXQlFjxkNJ6fO2eUP5jMqXvLdC", // admin123
		Nickname: "管理员",
		Email:    utils.StringPtr("admin@example.com"),
		Status:   1,
	}
	db.Create(&adminUser)

	// Assign admin role to admin user
	db.Create(&model.UserRole{UserID: adminUser.ID, RoleID: adminRole.ID})

	logger.Info("✅默认数据初始化完成")
}
