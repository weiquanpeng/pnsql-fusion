package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohongshu/PnSql/server/api/v1"
)

type Login struct{}

func (s *Login) InitLoginRouter(Router *gin.RouterGroup) {
	loginRouter := Router.Group("login")
	loginApi := v1.ApiGroupApp.SystemApiGroup.Login
	{
		loginRouter.GET("get", loginApi.Login)
		loginRouter.POST("post", loginApi.ToLogin)
	}
}
