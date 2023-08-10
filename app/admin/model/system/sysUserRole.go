package system

import (
	"ruoyi-go/app/admin/model/tools"
	"ruoyi-go/utils"
)

// SysUserRole model：数据库字段
type SysUserRole struct {
	UserId int    `json:"userId" gorm:"column:user_id;primaryKey"` //表示主键
	RoleId string `json:"roleId" gorm:"role_id"`
}

// TableName 指定数据库表名称
func (SysUserRole) TableName() string {
	return "sys_user_role"
}

// 分页查询
func SelectSysUserRoleList(params tools.SearchTableDataParam, isPage bool) tools.TableDataInfo {
	var pageNum = params.PageNum
	var pageSize = params.PageSize
	sysUserRole := params.Other.(SysUserRole)

	offset := (pageNum - 1) * pageSize
	var total int64
	var rows []SysUserRole

	var db = utils.MysqlDb.Model(&SysUserRole{}).
		Joins("left join sys_dept d on d.dept_id = dept_id").
		Select("*, d.dept_name, d.leader")

	db.Where("del_flag = '0'")
	var userId = sysUserRole.UserId
	if userId != 0 {
		db.Where("user_id = ?", userId)
	}
	if err := db.Count(&total).Error; err != nil {
		return tools.Fail()
	}
	if isPage {
		if err := db.Limit(pageSize).Offset(offset).Find(&rows).Error; err != nil {
			return tools.Fail()
		}
	} else {
		if err := db.Find(&rows).Error; err != nil {
			return tools.Fail()
		}
	}

	if rows == nil {
		return tools.Fail()
	} else {
		return tools.Success(rows, total)
	}
}
