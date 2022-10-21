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

type VideoSpecialFileTransApi struct {
}

var videoSpecialFileTransService = service.ServiceGroupApp.CloudServiceGroup.VideoSpecialFileTransService

// CreateVideoSpecialFileTrans 创建VideoSpecialFileTrans
// @Tags VideoSpecialFileTrans
// @Summary 创建VideoSpecialFileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoSpecialFileTrans true "创建VideoSpecialFileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoSpecialFileTrans/createVideoSpecialFileTrans [post]
func (videoSpecialFileTransApi *VideoSpecialFileTransApi) CreateVideoSpecialFileTrans(c *gin.Context) {
	var videoSpecialFileTrans cloud.VideoSpecialFileTrans
	err := c.ShouldBindJSON(&videoSpecialFileTrans)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := videoSpecialFileTransService.CreateVideoSpecialFileTrans(videoSpecialFileTrans); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteVideoSpecialFileTrans 删除VideoSpecialFileTrans
// @Tags VideoSpecialFileTrans
// @Summary 删除VideoSpecialFileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoSpecialFileTrans true "删除VideoSpecialFileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoSpecialFileTrans/deleteVideoSpecialFileTrans [delete]
func (videoSpecialFileTransApi *VideoSpecialFileTransApi) DeleteVideoSpecialFileTrans(c *gin.Context) {
	var videoSpecialFileTrans cloud.VideoSpecialFileTrans
	err := c.ShouldBindJSON(&videoSpecialFileTrans)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := videoSpecialFileTransService.DeleteVideoSpecialFileTrans(videoSpecialFileTrans); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteVideoSpecialFileTransByIds 批量删除VideoSpecialFileTrans
// @Tags VideoSpecialFileTrans
// @Summary 批量删除VideoSpecialFileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VideoSpecialFileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /videoSpecialFileTrans/deleteVideoSpecialFileTransByIds [delete]
func (videoSpecialFileTransApi *VideoSpecialFileTransApi) DeleteVideoSpecialFileTransByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := videoSpecialFileTransService.DeleteVideoSpecialFileTransByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateVideoSpecialFileTrans 更新VideoSpecialFileTrans
// @Tags VideoSpecialFileTrans
// @Summary 更新VideoSpecialFileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoSpecialFileTrans true "更新VideoSpecialFileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoSpecialFileTrans/updateVideoSpecialFileTrans [put]
func (videoSpecialFileTransApi *VideoSpecialFileTransApi) UpdateVideoSpecialFileTrans(c *gin.Context) {
	var videoSpecialFileTrans cloud.VideoSpecialFileTrans
	err := c.ShouldBindJSON(&videoSpecialFileTrans)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := videoSpecialFileTransService.UpdateVideoSpecialFileTrans(videoSpecialFileTrans); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindVideoSpecialFileTrans 用id查询VideoSpecialFileTrans
// @Tags VideoSpecialFileTrans
// @Summary 用id查询VideoSpecialFileTrans
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloud.VideoSpecialFileTrans true "用id查询VideoSpecialFileTrans"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoSpecialFileTrans/findVideoSpecialFileTrans [get]
func (videoSpecialFileTransApi *VideoSpecialFileTransApi) FindVideoSpecialFileTrans(c *gin.Context) {
	var videoSpecialFileTrans cloud.VideoSpecialFileTrans
	err := c.ShouldBindQuery(&videoSpecialFileTrans)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if revideoSpecialFileTrans, err := videoSpecialFileTransService.GetVideoSpecialFileTrans(videoSpecialFileTrans.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"revideoSpecialFileTrans": revideoSpecialFileTrans}, c)
	}
}

// GetVideoSpecialFileTransList 分页获取VideoSpecialFileTrans列表
// @Tags VideoSpecialFileTrans
// @Summary 分页获取VideoSpecialFileTrans列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloudReq.VideoSpecialFileTransSearch true "分页获取VideoSpecialFileTrans列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoSpecialFileTrans/getVideoSpecialFileTransList [get]
func (videoSpecialFileTransApi *VideoSpecialFileTransApi) GetVideoSpecialFileTransList(c *gin.Context) {
	var pageInfo cloudReq.VideoSpecialFileTransSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := videoSpecialFileTransService.GetVideoSpecialFileTransInfoList(pageInfo); err != nil {
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
