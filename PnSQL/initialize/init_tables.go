package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/xiaohongshu/PnSql/server/dao/model/example"
	"github.com/xiaohongshu/PnSql/server/dao/model/system"
	"github.com/xiaohongshu/PnSql/server/global"
	"os"
)

func RegisterTables() {
	db := global.PVA_DB
	err := db.AutoMigrate(
		// 系统模块表
		system.Jwt{},
		// 业务模块表
		example.SysUser{},
		example.SysRoles{},
		example.SysMenus{},
		example.SysUserMenus{},
	)
	if err != nil {
		fmt.Println("register table err:", err)
		os.Exit(0)
	}

	//初始化数据
	err = insertDefaultData()
	if err != nil {
		fmt.Println("insert default data err:", err)
		os.Exit(0)
	}
}

func insertDefaultData() error {
	//初始化 admin 用户
	var user_count int64
	global.PVA_DB.Model(&example.SysUser{}).Where("account = ?", "admin").Count(&user_count)
	if user_count == 0 {
		admin := example.SysUser{
			Account:  "admin",
			Password: "5c01e9c4dd9393e65ae259d39df61e8c",
			Enable:   1,
			Roles:    json.RawMessage(`[1, 2, 3]`),
		}
		result := global.PVA_DB.Create(&admin)
		if result.Error != nil {
			return result.Error
		}
	}

	//初始化权限表
	defaultRoles := []example.SysRoles{
		{RoleName: "超级用户", RoleCode: "1"},
		{RoleName: "管理用户", RoleCode: "2"},
		{RoleName: "业务用户", RoleCode: "3"},
	}
	defaultMenusRoles := []example.SysUserMenus{
		{UserID: 1, MenusID: 1}, {UserID: 1, MenusID: 10}, {UserID: 1, MenusID: 11}, {UserID: 1, MenusID: 12}, {UserID: 1, MenusID: 21}, {UserID: 1, MenusID: 22}, {UserID: 1, MenusID: 23}, {UserID: 1, MenusID: 24}, {UserID: 1, MenusID: 25},
		{UserID: 1, MenusID: 30}, {UserID: 1, MenusID: 31}, {UserID: 1, MenusID: 32}, {UserID: 1, MenusID: 40}, {UserID: 1, MenusID: 41},
		{UserID: 2, MenusID: 1}, {UserID: 2, MenusID: 10}, {UserID: 2, MenusID: 11}, {UserID: 2, MenusID: 12}, {UserID: 2, MenusID: 21}, {UserID: 2, MenusID: 22}, {UserID: 2, MenusID: 23},
		{UserID: 3, MenusID: 1}, {UserID: 3, MenusID: 10}, {UserID: 3, MenusID: 24}, {UserID: 3, MenusID: 25}, {UserID: 3, MenusID: 30}, {UserID: 3, MenusID: 31}, {UserID: 3, MenusID: 32}, {UserID: 3, MenusID: 33},
		{UserID: 4, MenusID: 1},
	}
	for _, role := range defaultRoles {
		var count int64
		global.PVA_DB.Model(&example.SysRoles{}).Where("role_name = ?", role.RoleName).Count(&count)
		if count == 0 {
			result := global.PVA_DB.Create(&role)
			if result.Error != nil {
				return result.Error
			}
		}
	}
	var user_menus_count int64
	global.PVA_DB.Model(&example.SysUserMenus{}).Count(&user_menus_count)
	if user_menus_count == 0 {
		result := global.PVA_DB.Create(&defaultMenusRoles)
		if result.Error != nil {
			return result.Error
		}
	}

	//初始化菜单表
	var menu_count int64
	global.PVA_DB.Model(&example.SysMenus{}).Where("name = ?", "主页").Count(&menu_count)
	if menu_count == 0 {
		defaultMenu := []example.SysMenus{
			{
				Name:     "主页",
				Path:     "/index",
				Icon:     "operation",
				Children: json.RawMessage(`[{"id": 1, "name":"大盘", "path": "/index/dashboard"}]`),
			},
			{
				Name:     "工单管理",
				Path:     "/task",
				Icon:     "SwitchFilled",
				Children: json.RawMessage(`[{"id": 10, "name":"我的工单", "path": "/db/mine"},{"id": 11, "name":"工单审批", "path": "/db/approve"},{"id": 12, "name":"所有工单", "path": "/db/all"}]`),
			},
			{
				Name:     "用户",
				Path:     "/sys",
				Icon:     "user",
				Children: json.RawMessage(`[{"id": 21, "name":"用户管理", "path": "/sys/user"},{"id": 22, "name":"用户权限", "path": "/sys/roles"}]`),
			},
			{
				Name:     "应用",
				Path:     "/app",
				Icon:     "menu",
				Children: json.RawMessage(`[{"id": 23, "name":"应用用户", "path": "/app/user"}]`),
			},
			{
				Name:     "测试",
				Path:     "/sit",
				Icon:     "odometer",
				Children: json.RawMessage(`[{"id": 24, "name":"测试环境", "path": "/sit/user"},{"id": 25, "name":"新测试环境", "path": "/sit/sit2"}]`),
			},
			{
				Name:     "数据库",
				Path:     "/db",
				Icon:     "list",
				Children: json.RawMessage(`[{"id": 30, "name":"MySQL", "path": "/db/mysql"},{"id": 31, "name":"MongoDB", "path": "/db/mongodb"},{"id": 32, "name":"Redis", "path": "/db/redis"},{"id": 33, "name":"Elasticsearch", "path": "/db/elasticsearch"}]`),
			},
			{
				Name:     "量化",
				Path:     "/quants",
				Icon:     "setting",
				Children: json.RawMessage(`[{"id": 40, "name":"量化选股", "path": "/quants/all"},{"id": 41, "name":"量化预测", "path": "/quants/traders"}]`),
			},
		}
		result := global.PVA_DB.Create(&defaultMenu)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
