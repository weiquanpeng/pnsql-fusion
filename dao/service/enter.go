package service

import (
	"github.com/xiaohongshu/PnSql/server/dao/service/example"
	"github.com/xiaohongshu/PnSql/server/dao/service/system"
)

type Enter struct {
	SystemServer  system.SysService
	ExampleServer example.ExampleService
}

var GroupApp = new(Enter)
