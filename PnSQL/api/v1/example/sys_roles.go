package example

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohongshu/PnSql/server/common/response"
)

type SysRolesApi struct{}

func (api *SysRolesApi) ListSysRoles(c *gin.Context) {
	sysUser, err := exampleService.SysRolesService.ListSysRoles()
	if err != nil {
		response.FailWithMessage("查询权限失败", c)
		return
	}
	response.OkWithDetailed(sysUser, "获取角色成功", c)
}
