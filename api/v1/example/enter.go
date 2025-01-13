package example

import (
	"github.com/xiaohongshu/PnSql/server/dao/service"
)

type ExaApi struct {
	SysUserApi
	SysRolesApi
	SysMenusApi
	TaskConfigApi
	SubTaskConfigApi
}

var (
	exampleService = service.GroupApp.ExampleServer
)
