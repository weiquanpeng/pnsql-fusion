package example

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohongshu/PnSql/server/common/response"
)

type SysMenusApi struct{}

func (api *SysMenusApi) ListMenus(c *gin.Context) {
	sysMenus, err := exampleService.SysMenusService.ListSysMenus()
	if err != nil {
		response.FailWithMessage("查询菜单失败", c)
		return
	}
	response.OkWithDetailed(sysMenus, "获取菜单成功", c)
}
