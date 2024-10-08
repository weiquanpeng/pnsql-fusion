package example

import (
	"github.com/xiaohongshu/PnSql/server/common/request"
	"github.com/xiaohongshu/PnSql/server/dao/model/example"
	"github.com/xiaohongshu/PnSql/server/dao/service/commons"
	"github.com/xiaohongshu/PnSql/server/global"
)

type SysUserService struct {
	commons.BaseService[example.SysUser]
}

// 用于登录
func (service *SysUserService) GetSysUsersAccount(account string) (user *example.SysUser, err error) {
	err = global.PVA_DB.Where("account = ?", account).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 根据id查询信息
func (service *SysUserService) GetSysUserByID(id uint) (user *example.SysUser, err error) {
	err = global.PVA_DB.Where("id = ?", id).First(&user).Error
	return
}

// 添加用户
func (service *SysUserService) CreateSysUsersAccount(user *example.SysUser) (err error) {
	err = global.PVA_DB.Create(&user).Error
	return err
}

// 修改用户
func (service *SysUserService) UpdateSysUsersAccount(user *example.SysUser) (err error) {
	err = global.PVA_DB.Model(&user).Updates(&user).Error //会自动根据主键 id 更新
	return err
}

// 根据 id 修改字段
func (service *SysUserService) UpdateUserField(id uint, fields map[string]interface{}) (err error) {
	var user example.SysUser
	err = global.PVA_DB.Model(&user).Where("id = ?", id).Updates(fields).Error
	return err
}

// 删除用户
func (service *SysUserService) DelSysUsersAccount(id uint) (err error) {
	var user example.SysUser
	err = global.PVA_DB.Where("id = ?", id).Delete(&user).Error
	return err
}

// 批量删除用户
func (service *SysUserService) DeleteSysUsersAccount(users []example.SysUser) (err error) {
	err = global.PVA_DB.Delete(&users).Error
	return err
}

// 分页查询
func (service *SysUserService) LoadSysUserPage(info request.PageInfo) (list interface{}, total int64, err error) {
	// 获取分页的参数信息
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 准备查询那个数据库表
	db := global.PVA_DB.Model(&example.SysUser{})

	// 准备切片帖子数组
	var sysUserList []example.SysUser

	// 加条件
	if info.Keyword != "" {
		db = db.Where("(account like ? or phone like ? )", "%"+info.Keyword+"%", "%"+info.Keyword+"%")
	}

	// 排序默时间降序降序
	db = db.Order("created_at desc")

	// 查询中枢
	err = db.Count(&total).Error
	if err != nil {
		return sysUserList, total, err
	} else {
		// 执行查询
		err = db.Limit(limit).Offset(offset).Find(&sysUserList).Error
	}

	// 结果返回
	return sysUserList, total, err
}
