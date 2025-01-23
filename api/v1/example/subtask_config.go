package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaohongshu/PnSql/server/common/response"
	"github.com/xiaohongshu/PnSql/server/global"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type SubTaskConfigApi struct{}

func (api *SubTaskConfigApi) GetSubTaskConfigData(c *gin.Context) {
	id := c.Query("id")
	parseUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.FailWithMessage("id 参数格式错误", c)
		return
	}
	// 获取子任务配置数据
	subTaskData, err := exampleService.SubTaskConfigService.GetSubTaskConfigData(uint(parseUint))
	if err != nil {
		response.FailWithMessage("工单详情获取失败: "+err.Error(), c)
		return
	}
	// 读取日志文件内容
	logFilePath := filepath.Join(global.P_cfg.System.Job_log, fmt.Sprintf("job_%s.log", id))
	logContent, err := os.ReadFile(logFilePath)
	if err != nil {
		global.Logger.Error("读取日志文件失败: " + err.Error())
		response.FailWithData(200, map[string]interface{}{
			"data": subTaskData, // 子任务配置数据
			"logs": "日志文件读取失败",  // 日志内容为空
		}, "读取日志文件失败: "+err.Error(), c)
		return
	}

	// 将日志内容按行拆分
	logLines := strings.Split(string(logContent), "\n")

	// 返回子任务配置数据和日志内容
	response.OkWithDetailed(map[string]interface{}{
		"data": subTaskData, // 子任务配置数据
		"logs": logLines,    // 日志内容
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

func (api *SubTaskConfigApi) UptSubTaskConfigData(c *gin.Context) {
	type Request struct {
		ID int64 `json:"id"` // 子工单 ID
	}

	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("请求参数错误", c)
		return
	}

	if err := exampleService.SubTaskConfigService.UptSubTaskConfigStatusByID(req.ID); err != nil {
		response.FailWithMessage("子工单完成失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("子工单状态修改成功", c)
}
