package example

import (
	"github.com/xiaohongshu/PnSql/server/dao/model/example"
	"github.com/xiaohongshu/PnSql/server/dao/service/commons"
	"github.com/xiaohongshu/PnSql/server/global"
)

type SysRolesService struct {
	commons.BaseService[example.SysRoles]
}

// 查询角色
func (service *SysRolesService) ListSysRoles() (roles []*example.SysRoles, err error) {
	err = global.PVA_DB.Order("id asc").Find(&roles).Error
	return
}
