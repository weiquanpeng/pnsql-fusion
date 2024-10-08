package system

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohongshu/PnSql/server/common/response"
	"github.com/xiaohongshu/PnSql/server/dao/model/system"
	"github.com/xiaohongshu/PnSql/server/dao/service"
)

type Logout struct{}

var jwtService = service.GroupApp.SystemServer.Jwt

func (l *Logout) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		response.FailWithMessage("请求未携带 token", c)
		return
	}

	err := jwtService.CreateJwt(&system.Jwt{Jwt: token})
	if err != nil {
		response.FailWithMessage("token限制失败", c)
		return
	}
	response.OkWithMessage("token已入黑名单", c)
}
