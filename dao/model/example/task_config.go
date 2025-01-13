package example

import (
	"encoding/json"
	"github.com/xiaohongshu/PnSql/server/global"
)

type TaskConfig struct {
	global.PvaModel
	Title        string          `json:"title" gorm:"comment:任务名"`
	Owner        string          `json:"owner"  gorm:"comment:提交人"`
	Status       string          `json:"status"  gorm:"index;comment:工单状态"`
	Parameter    json.RawMessage `json:"parameter" gorm:"type:json;comment:配置参数"`
	TaskDescribe string          `json:"describe"  gorm:"comment:工单描述"`
}

func (s *TaskConfig) TableName() string {
	return "task_config"
}
