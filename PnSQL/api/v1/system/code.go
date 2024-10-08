package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/xiaohongshu/PnSql/server/common/response"
)

type CodeApi struct{}

type codeJsonBody struct {
	Id   string `form:"id"`
	Code string `form:"code"`
}

// 创建验证码
func (codeApi *CodeApi) CreateCaptcha(c *gin.Context) {
	// 生成验证码类型
	var store = base64Captcha.DefaultMemStore
	digit := &base64Captcha.DriverDigit{Height: 80, Width: 240, Length: 6, MaxSkew: 0.8, DotCount: 120}
	captcha := base64Captcha.NewCaptcha(digit, store)
	id, baseURL, _, err := captcha.Generate()
	if err != nil {
		response.FailWithMessage("验证生成错误", c)
		return
	}
	response.OkWithData(map[string]any{"id": id, "baseURL": baseURL}, c)
}

func (codeApi *CodeApi) VerifyCaptcha(c *gin.Context) {
	newCode := codeJsonBody{}
	err := c.ShouldBind(&newCode)
	if err != nil {
		response.FailWithMessage("参数绑定失败", c)
		return
	}
	verify := store.Verify(newCode.Id, newCode.Code, true)
	if verify {
		response.Ok(c)
	} else {
		response.FailWithMessage("验证失败", c)
	}

}
