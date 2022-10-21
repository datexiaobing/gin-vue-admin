package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type FileDownloadApi struct {
}

var fileDownloadService = service.ServiceGroupApp.CloudServiceGroup.FileDownloadService


// CreateFileDownload 创建FileDownload
// @Tags FileDownload
// @Summary 创建FileDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.FileDownload true "创建FileDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileDownload/createFileDownload [post]
func (fileDownloadApi *FileDownloadApi) CreateFileDownload(c *gin.Context) {
	var fileDownload cloud.FileDownload
	err := c.ShouldBindJSON(&fileDownload)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    fileDownload.CreatedBy = utils.GetUserID(c)
	if err := fileDownloadService.CreateFileDownload(fileDownload); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFileDownload 删除FileDownload
// @Tags FileDownload
// @Summary 删除FileDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.FileDownload true "删除FileDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fileDownload/deleteFileDownload [delete]
func (fileDownloadApi *FileDownloadApi) DeleteFileDownload(c *gin.Context) {
	var fileDownload cloud.FileDownload
	err := c.ShouldBindJSON(&fileDownload)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    fileDownload.DeletedBy = utils.GetUserID(c)
	if err := fileDownloadService.DeleteFileDownload(fileDownload); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFileDownloadByIds 批量删除FileDownload
// @Tags FileDownload
// @Summary 批量删除FileDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FileDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fileDownload/deleteFileDownloadByIds [delete]
func (fileDownloadApi *FileDownloadApi) DeleteFileDownloadByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := fileDownloadService.DeleteFileDownloadByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFileDownload 更新FileDownload
// @Tags FileDownload
// @Summary 更新FileDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.FileDownload true "更新FileDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fileDownload/updateFileDownload [put]
func (fileDownloadApi *FileDownloadApi) UpdateFileDownload(c *gin.Context) {
	var fileDownload cloud.FileDownload
	err := c.ShouldBindJSON(&fileDownload)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    fileDownload.UpdatedBy = utils.GetUserID(c)
	if err := fileDownloadService.UpdateFileDownload(fileDownload); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFileDownload 用id查询FileDownload
// @Tags FileDownload
// @Summary 用id查询FileDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloud.FileDownload true "用id查询FileDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fileDownload/findFileDownload [get]
func (fileDownloadApi *FileDownloadApi) FindFileDownload(c *gin.Context) {
	var fileDownload cloud.FileDownload
	err := c.ShouldBindQuery(&fileDownload)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refileDownload, err := fileDownloadService.GetFileDownload(fileDownload.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refileDownload": refileDownload}, c)
	}
}

// GetFileDownloadList 分页获取FileDownload列表
// @Tags FileDownload
// @Summary 分页获取FileDownload列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloudReq.FileDownloadSearch true "分页获取FileDownload列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileDownload/getFileDownloadList [get]
func (fileDownloadApi *FileDownloadApi) GetFileDownloadList(c *gin.Context) {
	var pageInfo cloudReq.FileDownloadSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := fileDownloadService.GetFileDownloadInfoList(pageInfo); err != nil {
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
