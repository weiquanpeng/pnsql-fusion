package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaohongshu/PnSql/server/common/filter"
	"github.com/xiaohongshu/PnSql/server/common/jwt"
	"github.com/xiaohongshu/PnSql/server/global"
	"github.com/xiaohongshu/PnSql/server/router"
)

func InitRouter(args map[string]string) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(filter.Cors())

	systemRouter := router.AllRouter.SysGroup
	exampleRouter := router.AllRouter.ExaGroup

	// 公共路由组
	PublicGroup := r.Group("/api/v1/public")
	{
		systemRouter.InitLoginRouter(PublicGroup)
		exampleRouter.InitSysMenusRouter(PublicGroup)
	}

	// 权限路由组
	PrivateGroup := r.Group("/api/v1/")
	PrivateGroup.Use(jwt.JwtAuth())
	{
		systemRouter.InitLogoutRouter(PrivateGroup)
		exampleRouter.InitSysUserRouter(PrivateGroup)
		exampleRouter.InitSysRolesRouter(PrivateGroup)
	}

	stPort := global.P_cfg.System.Port
	if stPort == "" {
		stPort = "2379"
	}
	err := r.Run(fmt.Sprintf(":%s", stPort))

	//获取命令行 env
	var port = args["port"]
	if port == "" {
		port = global.P_cfg.System.Port
		if port == "" {
			port = "2379"
		}
	}

	err = r.Run(fmt.Sprintf(":%s", stPort))
	if err != nil {
		panic(fmt.Sprintf("start service error: %s", err.Error()))
	}
}
