package system

import (
	"errors"
	"github.com/xiaohongshu/PnSql/server/dao/model/system"
	"github.com/xiaohongshu/PnSql/server/global"
	"gorm.io/gorm"
)

type Jwt struct{}

func (a *Jwt) CreateJwt(jwt *system.Jwt) (err error) {
	err = global.PVA_DB.Create(jwt).Error
	return err
}

func (a *Jwt) GetJwt(jwt string) bool {
	var line system.Jwt
	err := global.PVA_DB.Where("jwt = ?", jwt).First(&line).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}
