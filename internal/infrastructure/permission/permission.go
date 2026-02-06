package permission

const (
	// 用户管理权限
	PermUserCreate      = "user:create"
	PermUserRead        = "user:read"
	PermUserUpdate      = "user:update"
	PermUserDelete      = "user:delete"
	PermUserList        = "user:list"
	PermUserAssignRoles = "user:assign-roles"

	// 角色管理权限
	PermRoleCreate            = "role:create"
	PermRoleRead              = "role:read"
	PermRoleUpdate            = "role:update"
	PermRoleDelete            = "role:delete"
	PermRoleList              = "role:list"
	PermRoleAssignPermissions = "role:assign-permissions"

	// 权限管理权限
	PermPermissionCreate = "permission:create"
	PermPermissionRead   = "permission:read"
	PermPermissionUpdate = "permission:update"
	PermPermissionDelete = "permission:delete"
	PermPermissionList   = "permission:list"
)
