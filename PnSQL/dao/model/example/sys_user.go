package example

import (
	"encoding/json"
	"github.com/xiaohongshu/PnSql/server/global"
)

type SysUser struct {
	global.PvaModel
	Account  string          `json:"account" gorm:"index;comment:用户登录名"`
	Password string          `json:"password"  gorm:"comment:用户登录密码"`
	Roles    json.RawMessage `json:"roles" gorm:"type:json;comment:用户权限"`
	Phone    string          `json:"phone"  gorm:"comment:用户电话"`
	Email    string          `json:"email"  gorm:"comment:用户邮箱"`
	Enable   int             `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
}

func (s *SysUser) TableName() string {
	return "sys_users"
}
