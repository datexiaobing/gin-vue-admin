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

type VideoListApi struct {
}

var videoListService = service.ServiceGroupApp.CloudServiceGroup.VideoListService


// CreateVideoList 创建VideoList
// @Tags VideoList
// @Summary 创建VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoList true "创建VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoList/createVideoList [post]
func (videoListApi *VideoListApi) CreateVideoList(c *gin.Context) {
	var videoList cloud.VideoList
	err := c.ShouldBindJSON(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoList.CreatedBy = utils.GetUserID(c)
	if err := videoListService.CreateVideoList(videoList); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteVideoList 删除VideoList
// @Tags VideoList
// @Summary 删除VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoList true "删除VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoList/deleteVideoList [delete]
func (videoListApi *VideoListApi) DeleteVideoList(c *gin.Context) {
	var videoList cloud.VideoList
	err := c.ShouldBindJSON(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoList.DeletedBy = utils.GetUserID(c)
	if err := videoListService.DeleteVideoList(videoList); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteVideoListByIds 批量删除VideoList
// @Tags VideoList
// @Summary 批量删除VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /videoList/deleteVideoListByIds [delete]
func (videoListApi *VideoListApi) DeleteVideoListByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := videoListService.DeleteVideoListByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateVideoList 更新VideoList
// @Tags VideoList
// @Summary 更新VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoList true "更新VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoList/updateVideoList [put]
func (videoListApi *VideoListApi) UpdateVideoList(c *gin.Context) {
	var videoList cloud.VideoList
	err := c.ShouldBindJSON(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoList.UpdatedBy = utils.GetUserID(c)
	if err := videoListService.UpdateVideoList(videoList); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindVideoList 用id查询VideoList
// @Tags VideoList
// @Summary 用id查询VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloud.VideoList true "用id查询VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoList/findVideoList [get]
func (videoListApi *VideoListApi) FindVideoList(c *gin.Context) {
	var videoList cloud.VideoList
	err := c.ShouldBindQuery(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if revideoList, err := videoListService.GetVideoList(videoList.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"revideoList": revideoList}, c)
	}
}

// GetVideoListList 分页获取VideoList列表
// @Tags VideoList
// @Summary 分页获取VideoList列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloudReq.VideoListSearch true "分页获取VideoList列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoList/getVideoListList [get]
func (videoListApi *VideoListApi) GetVideoListList(c *gin.Context) {
	var pageInfo cloudReq.VideoListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := videoListService.GetVideoListInfoList(pageInfo); err != nil {
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
