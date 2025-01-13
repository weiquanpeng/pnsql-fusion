package denominator

import (
	"github.com/xiaohongshu/PnSql/server/common/request"
	"github.com/xiaohongshu/PnSql/server/global"
)

type BaseService[T any] struct{}

// 根据id查询信息
func (service *BaseService[T]) GetByID(id uint) (user *T, err error) {
	err = global.PVA_DB.Where("id = ?", id).First(&user).Error
	return
}

// 添加
func (service *BaseService[T]) CreateMould(user *T) (err error) {
	err = global.PVA_DB.Create(&user).Error
	return err
}

// 修改
func (service *BaseService[T]) UpdateMould(user *T) (err error) {
	err = global.PVA_DB.Model(&user).Updates(&user).Error //会自动根据主键 id 更新
	return err
}

// 根据 id 修改字段
func (service *BaseService[T]) UpdateMouldField(id uint, fields map[string]interface{}) (err error) {
	var user T
	err = global.PVA_DB.Model(&user).Where("id = ?", id).Updates(fields).Error
	return err
}

// 删除
func (service *BaseService[T]) DelMould(id uint) (err error) {
	var user T
	err = global.PVA_DB.Where("id = ?", id).Delete(&user).Error
	return err
}

// 分页查询
func (service *BaseService[T]) LoadMouldPage(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	var data T
	db := global.PVA_DB.Model(&data)
	var sysUserList []T
	if info.Keyword != "" {
		db = db.Where("(account like ? or phone like ? )", "%"+info.Keyword+"%", "%"+info.Keyword+"%")
	}
	db = db.Order("created_at desc")
	err = db.Count(&total).Error
	if err != nil {
		return sysUserList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&sysUserList).Error
	}
	return sysUserList, total, err
}
