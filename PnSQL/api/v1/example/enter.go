package example

import (
	"github.com/xiaohongshu/PnSql/server/dao/service"
)

type ExaApi struct {
	SysUserApi
	SysRolesApi
	SysMenusApi
}

var (
	exampleService = service.GroupApp.ExampleServer
)
