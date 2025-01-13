package system

import (
	"github.com/xiaohongshu/PnSql/server/dao/model/system"
	"github.com/xiaohongshu/PnSql/server/global"
)

type Jwt struct{}

func (a *Jwt) CreateJwt(jwt *system.Jwt) (err error) {
	err = global.PVA_DB.Create(jwt).Error
	return err
}

func (a *Jwt) GetJwtCount(jwt string) (int64, error) {
	var count int64
	err := global.PVA_DB.Model(&system.Jwt{}).Where("jwt = ?", jwt).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
