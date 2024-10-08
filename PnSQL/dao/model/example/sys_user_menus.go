package example

type SysUserMenus struct {
	ID      int `gorm:"primary" json:"id"`             // 主键ID
	UserID  int `json:"user_id" gorm:"comment:用户 id"`  // 角色名
	MenusID int `json:"menus_id" gorm:"comment:菜单 id"` // 角色代号
}

func (s *SysUserMenus) TableName() string {
	return "sys_user_menus"
}
