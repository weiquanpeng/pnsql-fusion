package system

import "github.com/xiaohongshu/PnSql/server/global"

type Jwt struct {
	global.PvaModel
	Jwt string `gorm:"type:text;column:jwt"`
}

func (Jwt) TableName() string {
	return "jwt"
}
