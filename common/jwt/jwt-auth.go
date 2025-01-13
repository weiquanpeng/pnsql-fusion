package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohongshu/PnSql/server/common/response"
	"github.com/xiaohongshu/PnSql/server/dao/service"
)

var jwtService = service.GroupApp.SystemServer.Jwt

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.FailWithData(801, "", "请求未携带 token", c)
			c.Abort()
			return
		}
		judgmentJwt, err := jwtService.GetJwtCount(token)
		if judgmentJwt > 0 {
			response.FailWithData(801, "", "token 已被加入黑名单", c)
			c.Abort()
			return
		}
		myJwt := NewJWT()
		claims, err := myJwt.ParseToken(token)
		if err != nil {
			response.FailWithData(801, "", "token 失效", c)
			c.Abort()
			return
		}

		////token 续期
		//if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
		//	newTime, _ := utils.ParseDuration("7d")
		//	claims.ExpiresAt = time.Now().Add(newTime).Unix()
		//	newToken, _ := myJwt.CreateTokenByOldToken(token, *claims)
		//	c.Header("PanGu-token", newToken)
		//	err := jwtService.CreateJwt(&system.Jwt{Jwt: token})
		//	if err != nil {
		//		return
		//	}
		//}

		c.Set("claims", claims)
		c.Next()
	}
}
