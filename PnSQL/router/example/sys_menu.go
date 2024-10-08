package example

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohongshu/PnSql/server/api/v1"
)

type SysMenusRouter struct{}

func (s *SysMenusRouter) InitSysMenusRouter(Router *gin.RouterGroup) {
	sysMenusRouter := Router.Group("/menu")
	sysMenusApi := v1.ApiGroupApp.ExampleApiGroup.SysMenusApi
	{
		sysMenusRouter.GET("/list", sysMenusApi.ListMenus)
	}
}
