package example

import "time"

type ProcessSubtaskConfig struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`                                                     // 主键，自增
	Name         string    `gorm:"type:varchar(64);not null;uniqueIndex:udx_name"`                               // 任务名，唯一索引
	Type         string    `gorm:"type:varchar(64);not null;default:'job'"`                                      // 子任务类型: job, cronjob
	Description  string    `gorm:"type:varchar(255);not null;default:''"`                                        // 描述
	ExecType     string    `gorm:"type:varchar(64);not null;default:'bash'"`                                     // 操作类型: bash, ssh
	ScriptHost   string    `gorm:"type:varchar(64);not null;default:'localhost'"`                                // 脚本所在机器: localhost, ip
	Interpreter  string    `gorm:"type:varchar(64);not null;default:'python3'"`                                  // 解释器: python3, bash
	Script       string    `gorm:"type:varchar(128);not null;default:''"`                                        // 脚本路径
	IsControlled int       `gorm:"type:int;not null;default:0"`                                                  // 是否管控
	Owner        string    `gorm:"type:varchar(64);not null;default:'sikui'"`                                    // owner
	CreateTime   time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`                             // 创建时间
	UpdateTime   time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // 更新时间
}

// TableName 设置表名
func (ProcessSubtaskConfig) TableName() string {
	return "process_subtask_config"
}
