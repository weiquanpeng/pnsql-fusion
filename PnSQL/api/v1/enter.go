package v1

import (
	"github.com/xiaohongshu/PnSql/server/api/v1/example"
	"github.com/xiaohongshu/PnSql/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.SysApi
	ExampleApiGroup example.ExaApi
}

var ApiGroupApp = new(ApiGroup)
