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

type VideoSpecialApi struct {
}

var videoSpecialService = service.ServiceGroupApp.CloudServiceGroup.VideoSpecialService


// CreateVideoSpecial 创建VideoSpecial
// @Tags VideoSpecial
// @Summary 创建VideoSpecial
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoSpecial true "创建VideoSpecial"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoSpecial/createVideoSpecial [post]
func (videoSpecialApi *VideoSpecialApi) CreateVideoSpecial(c *gin.Context) {
	var videoSpecial cloud.VideoSpecial
	err := c.ShouldBindJSON(&videoSpecial)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoSpecial.CreatedBy = utils.GetUserID(c)
	if err := videoSpecialService.CreateVideoSpecial(videoSpecial); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteVideoSpecial 删除VideoSpecial
// @Tags VideoSpecial
// @Summary 删除VideoSpecial
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoSpecial true "删除VideoSpecial"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoSpecial/deleteVideoSpecial [delete]
func (videoSpecialApi *VideoSpecialApi) DeleteVideoSpecial(c *gin.Context) {
	var videoSpecial cloud.VideoSpecial
	err := c.ShouldBindJSON(&videoSpecial)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoSpecial.DeletedBy = utils.GetUserID(c)
	if err := videoSpecialService.DeleteVideoSpecial(videoSpecial); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteVideoSpecialByIds 批量删除VideoSpecial
// @Tags VideoSpecial
// @Summary 批量删除VideoSpecial
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VideoSpecial"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /videoSpecial/deleteVideoSpecialByIds [delete]
func (videoSpecialApi *VideoSpecialApi) DeleteVideoSpecialByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := videoSpecialService.DeleteVideoSpecialByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateVideoSpecial 更新VideoSpecial
// @Tags VideoSpecial
// @Summary 更新VideoSpecial
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoSpecial true "更新VideoSpecial"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoSpecial/updateVideoSpecial [put]
func (videoSpecialApi *VideoSpecialApi) UpdateVideoSpecial(c *gin.Context) {
	var videoSpecial cloud.VideoSpecial
	err := c.ShouldBindJSON(&videoSpecial)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    videoSpecial.UpdatedBy = utils.GetUserID(c)
	if err := videoSpecialService.UpdateVideoSpecial(videoSpecial); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindVideoSpecial 用id查询VideoSpecial
// @Tags VideoSpecial
// @Summary 用id查询VideoSpecial
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloud.VideoSpecial true "用id查询VideoSpecial"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoSpecial/findVideoSpecial [get]
func (videoSpecialApi *VideoSpecialApi) FindVideoSpecial(c *gin.Context) {
	var videoSpecial cloud.VideoSpecial
	err := c.ShouldBindQuery(&videoSpecial)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if revideoSpecial, err := videoSpecialService.GetVideoSpecial(videoSpecial.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"revideoSpecial": revideoSpecial}, c)
	}
}

// GetVideoSpecialList 分页获取VideoSpecial列表
// @Tags VideoSpecial
// @Summary 分页获取VideoSpecial列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloudReq.VideoSpecialSearch true "分页获取VideoSpecial列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoSpecial/getVideoSpecialList [get]
func (videoSpecialApi *VideoSpecialApi) GetVideoSpecialList(c *gin.Context) {
	var pageInfo cloudReq.VideoSpecialSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := videoSpecialService.GetVideoSpecialInfoList(pageInfo); err != nil {
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
