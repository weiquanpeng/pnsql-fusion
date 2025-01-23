package example

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohongshu/PnSql/server/api/v1"
)

type TaskConfigRouter struct{}

func (s *SysUserRouter) InitTaskConfigRouter(Router *gin.RouterGroup) {
	taskConfigRouter := Router.Group("/task")
	taskConfigApi := v1.ApiGroupApp.ExampleApiGroup.TaskConfigApi
	subTaskConfigApi := v1.ApiGroupApp.ExampleApiGroup.SubTaskConfigApi
	{
		taskConfigRouter.GET("/list", taskConfigApi.LoadTaskConfigPage)
		taskConfigRouter.GET("/subtask/list", subTaskConfigApi.GetSubTaskConfigData)
		taskConfigRouter.POST("/mine/list", taskConfigApi.LoadOwnerTaskConfigPage)
		taskConfigRouter.POST("/approve/list", subTaskConfigApi.LoadSubTaskConfigPage)
		taskConfigRouter.POST("/log/list", taskConfigApi.LoadOwnerTaskConfigPage)
		taskConfigRouter.POST("/approve/uptStatus", subTaskConfigApi.UptSubTaskConfigData)
	}
}
