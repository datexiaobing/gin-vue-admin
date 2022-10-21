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

type VideoCategoryApi struct {
}

var videoCategoryService = service.ServiceGroupApp.CloudServiceGroup.VideoCategoryService


// CreateVideoCategory 创建VideoCategory
// @Tags VideoCategory
// @Summary 创建VideoCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoCategory true "创建VideoCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoCategory/createVideoCategory [post]
func (videoCategoryApi *VideoCategoryApi) CreateVideoCategory(c *gin.Context) {
	var videoCategory cloud.VideoCategory
	err := c.ShouldBindJSON(&videoCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoCategory.CreatedBy = utils.GetUserID(c)
	if err := videoCategoryService.CreateVideoCategory(videoCategory); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteVideoCategory 删除VideoCategory
// @Tags VideoCategory
// @Summary 删除VideoCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoCategory true "删除VideoCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoCategory/deleteVideoCategory [delete]
func (videoCategoryApi *VideoCategoryApi) DeleteVideoCategory(c *gin.Context) {
	var videoCategory cloud.VideoCategory
	err := c.ShouldBindJSON(&videoCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoCategory.DeletedBy = utils.GetUserID(c)
	if err := videoCategoryService.DeleteVideoCategory(videoCategory); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteVideoCategoryByIds 批量删除VideoCategory
// @Tags VideoCategory
// @Summary 批量删除VideoCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VideoCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /videoCategory/deleteVideoCategoryByIds [delete]
func (videoCategoryApi *VideoCategoryApi) DeleteVideoCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := videoCategoryService.DeleteVideoCategoryByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateVideoCategory 更新VideoCategory
// @Tags VideoCategory
// @Summary 更新VideoCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoCategory true "更新VideoCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoCategory/updateVideoCategory [put]
func (videoCategoryApi *VideoCategoryApi) UpdateVideoCategory(c *gin.Context) {
	var videoCategory cloud.VideoCategory
	err := c.ShouldBindJSON(&videoCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoCategory.UpdatedBy = utils.GetUserID(c)
	if err := videoCategoryService.UpdateVideoCategory(videoCategory); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindVideoCategory 用id查询VideoCategory
// @Tags VideoCategory
// @Summary 用id查询VideoCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloud.VideoCategory true "用id查询VideoCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoCategory/findVideoCategory [get]
func (videoCategoryApi *VideoCategoryApi) FindVideoCategory(c *gin.Context) {
	var videoCategory cloud.VideoCategory
	err := c.ShouldBindQuery(&videoCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if revideoCategory, err := videoCategoryService.GetVideoCategory(videoCategory.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"revideoCategory": revideoCategory}, c)
	}
}

// GetVideoCategoryList 分页获取VideoCategory列表
// @Tags VideoCategory
// @Summary 分页获取VideoCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloudReq.VideoCategorySearch true "分页获取VideoCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoCategory/getVideoCategoryList [get]
func (videoCategoryApi *VideoCategoryApi) GetVideoCategoryList(c *gin.Context) {
	var pageInfo cloudReq.VideoCategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := videoCategoryService.GetVideoCategoryInfoList(pageInfo); err != nil {
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
