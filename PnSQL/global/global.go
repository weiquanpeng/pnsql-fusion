package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	PVA_DB *gorm.DB
	Logger *zap.SugaredLogger
)

type Pva struct {
	System System
	Mysql  Mysql
	Zap    Zap
}

var P_cfg = new(Pva)
