package example

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohongshu/PnSql/server/common/request"
	"github.com/xiaohongshu/PnSql/server/common/response"
	"github.com/xiaohongshu/PnSql/server/dao/model/example"
	"github.com/xiaohongshu/PnSql/server/global"
	"github.com/xiaohongshu/PnSql/server/utils"
	"strconv"
)

type SysUserApi struct{}

func (api *SysUserApi) GetById(c *gin.Context) {
	id := c.Param("id")
	parseUint, _ := strconv.ParseUint(id, 10, 64) //转换 Param id 为 int
	sysUser, err := exampleService.SysUserService.GetByID(uint(parseUint))
	if err != nil {
		global.Logger.Errorf("查询用户: %s 失败", id)
		response.FailWithMessage("查询用户失败", c)
		return
	}
	response.OkWithData(sysUser, c)
}

// AddSysUser 添加用户
func (api *SysUserApi) AddSysUser(c *gin.Context) {
	var sysUser example.SysUser
	if err := c.ShouldBindJSON(&sysUser); err != nil {
		response.FailWithMessage("用户参数不匹配", c)
		return
	}
	sysUser.Password = utils.Md5(sysUser.Password)
	err := exampleService.SysUserService.CreateMould(&sysUser)
	if err != nil {
		response.FailWithDecide(err.Error(), "用户创建失败", c)
		return
	}
	global.Logger.Info("用户(" + sysUser.Account + ")创建成功")
	response.OkWithMessage("用户创建成功", c)
}

// UptSysUser 修改用户
func (api *SysUserApi) UptSysUser(c *gin.Context) {
	var sysUser example.SysUser
	if err := c.ShouldBindJSON(&sysUser); err != nil {
		response.FailWithDecide(err.Error(), "参数不匹配", c)
		return
	}
	err := exampleService.SysUserService.UpdateMould(&sysUser)
	if err != nil {
		response.FailWithMessage("用户编辑失败", c)
		return
	}
	response.OkWithMessage("用户编辑成功", c)
}

// 根据 id 修改字段
func (api *SysUserApi) UptSysUserField(c *gin.Context) {
	var params struct {
		ID    uint                   `json:"id"`
		Filed map[string]interface{} `json:"field"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		response.FailWithDecide(err.Error(), "参数不匹配", c)
		return
	}
	if password, ok := params.Filed["password"].(string); ok {
		params.Filed["password"] = utils.Md5(password)
	}
	err := exampleService.SysUserService.UpdateMouldField(params.ID, params.Filed)
	if err != nil {
		response.FailWithDecide(err.Error(), "更新字段失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DelSysUser 删除用户
func (api *SysUserApi) DelSysUser(c *gin.Context) {
	id := c.Query("id")
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	err := exampleService.SysUserService.DelMould(uint(parseUint))
	if err != nil {
		response.FailWithDecide(err.Error(), "删除用户失败", c)
		return
	}
	response.OkWithMessage("删除用户成功", c)
}

// LoadSysUserPage 分页查询
func (api *SysUserApi) LoadSysUserPage(c *gin.Context) {
	var pageInfo request.PageInfo
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userPage, total, err := exampleService.SysUserService.LoadMouldPage(pageInfo)
	if err != nil {
		global.Logger.Error("分页查询失败")
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     userPage,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)

}
