package system

// SysUserPost model：数据库字段
type SysUserPost struct {
	PostId int `json:"postId"`
	UserId int `json:"userId"`
}

// TableName 指定数据库表名称
func (SysUserPost) TableName() string {
	return "sys_user_post"
}
