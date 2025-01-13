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
