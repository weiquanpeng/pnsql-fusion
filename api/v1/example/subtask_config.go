package example

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohongshu/PnSql/server/common/response"
	"strconv"
)

type SubTaskConfigApi struct{}

func (api *SubTaskConfigApi) GetSubTaskConfigData(c *gin.Context) {
	id := c.Query("id")
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	subTaskData, err := exampleService.SubTaskConfigService.GetSubTaskConfigData(uint(parseUint))

	if err != nil {
		response.FailWithMessage("工单详情获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: subTaskData,
	}, "获取成功", c)
}

func (api *SubTaskConfigApi) LoadSubTaskConfigPage(c *gin.Context) {
	type OwnerRequest struct {
		Owner string `json:"owner"`
	}
	var req OwnerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("工单列表请求参数错误", c)
		return
	}
	ids, err := exampleService.SubTaskConfigService.LoadSubTaskConfigPage(req.Owner)
	if err != nil {
		response.FailWithMessage("工单列表获取失败"+err.Error(), c)
		return
	}
	approvePage, err := exampleService.TaskConfigService.LoadIdTaskConfigPage(ids)
	if err != nil {
		response.FailWithMessage("我的审批获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: approvePage,
	}, "获取成功", c)
}
