package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileInputApi struct {
}

var fileInputService = service.ServiceGroupApp.CloudServiceGroup.FileInputService

// CreateFileInput 创建FileInput
// @Tags FileInput
// @Summary 创建FileInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.FileInput true "创建FileInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileInput/createFileInput [post]
func (fileInputApi *FileInputApi) CreateFileInput(c *gin.Context) {
	var fileInput cloud.FileInput
	err := c.ShouldBindJSON(&fileInput)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileInput.CreatedBy = utils.GetUserID(c)
	if err := fileInputService.CreateFileInput(fileInput); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFileInput 删除FileInput
// @Tags FileInput
// @Summary 删除FileInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.FileInput true "删除FileInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileInput/deleteFileInput [delete]
func (fileInputApi *FileInputApi) DeleteFileInput(c *gin.Context) {
	var fileInput cloud.FileInput
	err := c.ShouldBindJSON(&fileInput)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileInput.DeletedBy = utils.GetUserID(c)
	if err := fileInputService.DeleteFileInput(fileInput); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFileInputByIds 批量删除FileInput
// @Tags FileInput
// @Summary 批量删除FileInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FileInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fileInput/deleteFileInputByIds [delete]
func (fileInputApi *FileInputApi) DeleteFileInputByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := fileInputService.DeleteFileInputByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFileInput 更新FileInput
// @Tags FileInput
// @Summary 更新FileInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.FileInput true "更新FileInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileInput/updateFileInput [put]
func (fileInputApi *FileInputApi) UpdateFileInput(c *gin.Context) {
	var fileInput cloud.FileInput
	err := c.ShouldBindJSON(&fileInput)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileInput.UpdatedBy = utils.GetUserID(c)
	if err := fileInputService.UpdateFileInput(fileInput); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFileInput 用id查询FileInput
// @Tags FileInput
// @Summary 用id查询FileInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloud.FileInput true "用id查询FileInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileInput/findFileInput [get]
func (fileInputApi *FileInputApi) FindFileInput(c *gin.Context) {
	var fileInput cloud.FileInput
	err := c.ShouldBindQuery(&fileInput)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refileInput, err := fileInputService.GetFileInput(fileInput.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refileInput": refileInput}, c)
	}
}

// GetFileInputList 分页获取FileInput列表
// @Tags FileInput
// @Summary 分页获取FileInput列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloudReq.FileInputSearch true "分页获取FileInput列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileInput/getFileInputList [get]
func (fileInputApi *FileInputApi) GetFileInputList(c *gin.Context) {
	var pageInfo cloudReq.FileInputSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fileInputService.GetFileInputInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
