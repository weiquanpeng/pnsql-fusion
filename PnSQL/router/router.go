package router

import (
	"github.com/xiaohongshu/PnSql/server/router/example"
	"github.com/xiaohongshu/PnSql/server/router/system"
)

type EnterGroup struct {
	SysGroup system.SysGroup
	ExaGroup example.ExaGroup
}

var AllRouter = new(EnterGroup)
