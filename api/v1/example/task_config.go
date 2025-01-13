package example

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohongshu/PnSql/server/common/response"
	"github.com/xiaohongshu/PnSql/server/global"
)

type TaskConfigApi struct{}

// LoadTaskConfigPage 查询 1000行
func (api *TaskConfigApi) LoadTaskConfigPage(c *gin.Context) {
	userPage, total, err := exampleService.TaskConfigService.LoadTaskConfigPage()
	if err != nil {
		global.Logger.Error("主工单分页查询失败")
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:  userPage,
		Total: total,
	}, "获取成功", c)
}

func (api *TaskConfigApi) LoadOwnerTaskConfigPage(c *gin.Context) {
	type OwnerRequest struct {
		Owner string `json:"owner"`
	}
	var req OwnerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("我的工单请求参数错误", c)
		return
	}
	ownerUserPage, err := exampleService.TaskConfigService.LoadOwnerTaskConfigPage(req.Owner)
	if err != nil {
		global.Logger.Error("我的工单查询失败")
		response.FailWithMessage("我的工单获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: ownerUserPage,
	}, "获取成功", c)
}
