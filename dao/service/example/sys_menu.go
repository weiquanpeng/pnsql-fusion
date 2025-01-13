package example

import (
	"github.com/xiaohongshu/PnSql/server/dao/model/example"
	"github.com/xiaohongshu/PnSql/server/dao/service/denominator"
	"github.com/xiaohongshu/PnSql/server/global"
)

type SysMenusService struct {
	denominator.BaseService[example.SysMenus]
}

func (service *SysMenusService) ListSysMenus() (menus []*example.SysMenus, err error) {
	err = global.PVA_DB.Order("id asc").Find(&menus).Error
	return
}
