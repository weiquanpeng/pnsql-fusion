package initialize

import (
	"github.com/xiaohongshu/PnSql/server/core"
	"github.com/xiaohongshu/PnSql/server/global"
)

func InitConfig() {
	core.Viper()                      //初始化配置
	global.Logger = core.InitLogger() //初始化日志
	global.PVA_DB = core.GormMysql()  //初始化 DB
	RegisterTables()
}
