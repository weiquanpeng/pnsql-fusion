package example

import (
	"github.com/xiaohongshu/PnSql/server/dao/model/example"
	"github.com/xiaohongshu/PnSql/server/dao/service/denominator"
	"github.com/xiaohongshu/PnSql/server/global"
)

type TaskConfigService struct {
	denominator.BaseService[example.TaskConfig]
}

// 分页查询
func (service *TaskConfigService) LoadTaskConfigPage() (list interface{}, total int64, err error) {
	limit := 1000
	db := global.PVA_DB.Model(&example.TaskConfig{})
	var taskConfigList []example.TaskConfig
	db = db.Order("updated_at desc")
	err = db.Count(&total).Error
	if err != nil {
		return taskConfigList, total, err
	} else {
		err = db.Limit(limit).Find(&taskConfigList).Error
	}
	return taskConfigList, total, err
}

func (service *TaskConfigService) LoadOwnerTaskConfigPage(user string) (list interface{}, err error) {
	limit := 1000
	db := global.PVA_DB.Model(&example.TaskConfig{})
	var taskConfigList []example.TaskConfig
	db = db.Where("owner = ? AND status != ?", user, "close").Order("updated_at desc")
	err = db.Limit(limit).Find(&taskConfigList).Error
	return taskConfigList, err
}

func (service *TaskConfigService) LoadIdTaskConfigPage(ids []int64) (list []example.TaskConfig, err error) {
	limit := 1000
	db := global.PVA_DB.Model(&example.TaskConfig{})
	var taskConfigList []example.TaskConfig
	db = db.Where("id IN (?) AND status != ?", ids, "close").Order("updated_at desc")
	err = db.Limit(limit).Find(&taskConfigList).Error
	return taskConfigList, err
}
