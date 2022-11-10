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

type FileTransApi struct {
}

var fileTransService = service.ServiceGroupApp.CloudServiceGroup.FileTransService

// CreateFileTrans 创建FileTrans
// @Tags FileTrans
// @Summary 创建FileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.FileTrans true "创建FileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileTrans/createFileTrans [post]
func (fileTransApi *FileTransApi) CreateFileTrans(c *gin.Context) {
	var fileTrans request.TransVideoReq
	err := c.ShouldBindJSON(&fileTrans)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileTrans.FileTrans.CreatedBy = utils.GetUserID(c)
	if err := fileTransService.CreateFileTrans(fileTrans); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFileTrans 删除FileTrans
// @Tags FileTrans
// @Summary 删除FileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.FileTrans true "删除FileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileTrans/deleteFileTrans [delete]
func (fileTransApi *FileTransApi) DeleteFileTrans(c *gin.Context) {
	var fileTrans cloud.FileTrans
	err := c.ShouldBindJSON(&fileTrans)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileTrans.DeletedBy = utils.GetUserID(c)
	if err := fileTransService.DeleteFileTrans(fileTrans); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFileTransByIds 批量删除FileTrans
// @Tags FileTrans
// @Summary 批量删除FileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fileTrans/deleteFileTransByIds [delete]
func (fileTransApi *FileTransApi) DeleteFileTransByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := fileTransService.DeleteFileTransByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFileTrans 更新FileTrans
// @Tags FileTrans
// @Summary 更新FileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.FileTrans true "更新FileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileTrans/updateFileTrans [put]
func (fileTransApi *FileTransApi) UpdateFileTrans(c *gin.Context) {
	var fileTrans cloud.FileTrans
	err := c.ShouldBindJSON(&fileTrans)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileTrans.UpdatedBy = utils.GetUserID(c)
	if err := fileTransService.UpdateFileTrans(fileTrans); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFileTrans 用id查询FileTrans
// @Tags FileTrans
// @Summary 用id查询FileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloud.FileTrans true "用id查询FileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileTrans/findFileTrans [get]
func (fileTransApi *FileTransApi) FindFileTrans(c *gin.Context) {
	var fileTrans cloud.FileTrans
	err := c.ShouldBindQuery(&fileTrans)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refileTrans, err := fileTransService.GetFileTrans(fileTrans.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refileTrans": refileTrans}, c)
	}
}

// GetFileTransList 分页获取FileTrans列表
// @Tags FileTrans
// @Summary 分页获取FileTrans列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloudReq.FileTransSearch true "分页获取FileTrans列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileTrans/getFileTransList [get]
func (fileTransApi *FileTransApi) GetFileTransList(c *gin.Context) {
	var pageInfo cloudReq.FileTransSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fileTransService.GetFileTransInfoList(pageInfo); err != nil {
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

// get sgare list
func (fileTransApi *FileTransApi) GetShareList(c *gin.Context) {
	var share request.IdsShareReq
	err := c.ShouldBindJSON(&share)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := fileTransService.GetShareList(share); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}
