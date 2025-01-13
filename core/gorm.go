package core

import (
	"github.com/xiaohongshu/PnSql/server/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func GormMysql() *gorm.DB {
	m := global.P_cfg.Mysql
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         128,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.Discard, //设置日志级别不输出 sql
	}); err != nil {
		global.Logger.Error("mysql 连接报错", zap.Error(err))
		os.Exit(0)
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(50)
		sqlDB.SetMaxOpenConns(100)
		return db
	}
}
