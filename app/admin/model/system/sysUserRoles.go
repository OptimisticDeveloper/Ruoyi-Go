package system

// SysUserRoles model：数据库字段
type SysUserRoles struct {
	RoleId int `json:"roleId"`
	UserId int `json:"userId"`
}

type SysUserRolesParam struct {
	RoleId string `json:"roleId"`
	UserId int    `json:"userId"`
}

// TableName 指定数据库表名称
func (SysUserRoles) TableName() string {
	return "sys_user_role"
}
