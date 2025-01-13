package example

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohongshu/PnSql/server/api/v1"
)

type SysUserRouter struct{}

func (s *SysUserRouter) InitSysUserRouter(Router *gin.RouterGroup) {
	sysUserRouter := Router.Group("/sys")
	sysUserApi := v1.ApiGroupApp.ExampleApiGroup.SysUserApi
	{
		sysUserRouter.GET("/user/post/:id", sysUserApi.GetById)
		sysUserRouter.POST("/user/addUser", sysUserApi.AddSysUser)
		sysUserRouter.POST("/user/uptUser", sysUserApi.UptSysUser)
		sysUserRouter.POST("/user/uptFiled", sysUserApi.UptSysUserField)
		sysUserRouter.POST("/user/delUser", sysUserApi.DelSysUser)
		sysUserRouter.POST("/user/load", sysUserApi.LoadSysUserPage)
	}
}
