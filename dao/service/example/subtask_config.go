package example

import (
	"github.com/xiaohongshu/PnSql/server/dao/model/example"
	"github.com/xiaohongshu/PnSql/server/dao/service/denominator"
	"github.com/xiaohongshu/PnSql/server/global"
)

type SubTaskConfigService struct {
	denominator.BaseService[example.SubTaskConfig]
}

func (service *SubTaskConfigService) GetSubTaskConfigData(taskID uint) (list []example.SubTaskConfig, err error) {
	db := global.PVA_DB.Model(&example.SubTaskConfig{})
	var subTaskConfigList []example.SubTaskConfig

	err = db.Where("task_id = ?", taskID).
		Order("id").
		Find(&subTaskConfigList).Error

	return subTaskConfigList, err
}

func (service *SubTaskConfigService) LoadSubTaskConfigPage(user string) (list []int64, err error) {
	limit := 1000
	db := global.PVA_DB.Model(&example.SubTaskConfig{})
	err = db.Where("FIND_IN_SET(?, approve) AND status = ?", user, "todo").
		Limit(limit).
		Distinct("task_id").
		Pluck("task_id", &list).Error
	return list, err
}

func (service *SubTaskConfigService) UptSubTaskConfigStatusByID(id int64) error {
	// 使用全局的 global.PVA_DB 执行更新操作
	err := global.PVA_DB.Model(&example.SubTaskConfig{}).
		Where("id = ?", id).
		Update("status", "done").Error
	return err
}

func (service *SubTaskConfigService) UptSubTaskConfigStatusByTaskID(taskID int64) error {
	// 使用全局的 global.PVA_DB 执行更新操作
	err := global.PVA_DB.Model(&example.SubTaskConfig{}).
		Where("task_id = ? and status='tosplit'", taskID).
		Limit(1). // 限制更新一条记录
		Update("status", "todo").Error
	return err
}
