package example

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohongshu/PnSql/server/common/response"
	"github.com/xiaohongshu/PnSql/server/global"
)

type TaskLogApi struct{}

func (api *TaskLogApi) LoadTaskLogList(c *gin.Context) {
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
