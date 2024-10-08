package example

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohongshu/PnSql/server/api/v1"
)

type SysRolesRouter struct{}

func (s *SysUserRouter) InitSysRolesRouter(Router *gin.RouterGroup) {
	sysRolesRouter := Router.Group("/sys")
	sysRolesApi := v1.ApiGroupApp.ExampleApiGroup.SysRolesApi
	{
		sysRolesRouter.GET("/roles/list", sysRolesApi.ListSysRoles)
	}
}
