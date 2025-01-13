package example

import (
	"encoding/json"
	"github.com/xiaohongshu/PnSql/server/global"
)

type SubTaskConfig struct {
	global.PvaModel
	Title        string          `json:"title" gorm:"comment:任务名"`
	Owner        string          `json:"owner"  gorm:"index;comment:提交人"`
	Approve      string          `json:"approve"  gorm:"comment:审批人"`
	TaskID       uint            `json:"taskID" gorm:"index;comment:主任务ID"`
	Status       string          `json:"status"  gorm:"index;comment:工单状态"`
	Parameter    json.RawMessage `json:"parameter" gorm:"type:json;comment:配置参数"`
	TaskDescribe string          `json:"describe"  gorm:"comment:工单描述"`
}

func (s *SubTaskConfig) TableName() string {
	return "subtask_config"
}
