package system

import (
	"github.com/xiaohongshu/PnSql/server/dao/service"
)

type SysApi struct {
	Login
	Logout
}

var (
	systemLogin  = service.GroupApp.SystemServer
	exampleLogin = service.GroupApp.ExampleServer
)
