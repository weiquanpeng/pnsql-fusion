package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohongshu/PnSql/server/api/v1"
)

type Logout struct{}

func (s *Logout) InitLogoutRouter(Router *gin.RouterGroup) {
	logoutRouter := Router.Group("logout")
	logoutApi := v1.ApiGroupApp.SystemApiGroup.Logout
	{
		logoutRouter.POST("post", logoutApi.Logout)
	}
}
