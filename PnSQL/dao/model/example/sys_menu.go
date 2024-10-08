package example

import "encoding/json"

type SysMenus struct {
	ID       uint            `gorm:"primary" json:"id"`
	Name     string          `json:"name" gorm:"not null;comment:菜单名"`
	Path     string          `json:"path" gorm:"not null;unique;comment:路由节点"`
	Icon     string          `json:"icon" gorm:"comment:图标"`
	Children json.RawMessage `json:"children" gorm:"type:json;comment:子菜单"` // 存储子菜单为 JSON 列
}

func (s *SysMenus) TableName() string {
	return "sys_menus"
}
