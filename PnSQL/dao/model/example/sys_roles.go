package example

import "github.com/xiaohongshu/PnSql/server/global"

type SysRoles struct {
	global.PvaModel
	RoleName string `json:"roleName" gorm:"comment:角色名"`  // 角色名
	RoleCode string `json:"roleCode" gorm:"comment:角色代号"` // 角色代号
}

func (s *SysRoles) TableName() string {
	return "sys_roles"
}
