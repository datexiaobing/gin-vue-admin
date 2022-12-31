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

type CloudPostApi struct {
}

var cloudPostService = service.ServiceGroupApp.CloudServiceGroup.CloudPostService

// CreateCloudPost 创建CloudPost
// @Tags CloudPost
// @Summary 创建CloudPost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.CloudPost true "创建CloudPost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cloudPost/createCloudPost [post]
func (cloudPostApi *CloudPostApi) CreateCloudPost(c *gin.Context) {
	var cloudPost cloud.CloudPost
	err := c.ShouldBindJSON(&cloudPost)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cloudPost.CreatedBy = utils.GetUserID(c)
	if err := cloudPostService.CreateCloudPost(cloudPost); err != nil {
		global.GVA_LOG.Error("推送失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("推送成功", c)
	}
}

// DeleteCloudPost 删除CloudPost
// @Tags CloudPost
// @Summary 删除CloudPost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.CloudPost true "删除CloudPost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cloudPost/deleteCloudPost [delete]
func (cloudPostApi *CloudPostApi) DeleteCloudPost(c *gin.Context) {
	var cloudPost cloud.CloudPost
	err := c.ShouldBindJSON(&cloudPost)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cloudPost.DeletedBy = utils.GetUserID(c)
	if err := cloudPostService.DeleteCloudPost(cloudPost); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCloudPostByIds 批量删除CloudPost
// @Tags CloudPost
// @Summary 批量删除CloudPost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CloudPost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cloudPost/deleteCloudPostByIds [delete]
func (cloudPostApi *CloudPostApi) DeleteCloudPostByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := cloudPostService.DeleteCloudPostByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCloudPost 更新CloudPost
// @Tags CloudPost
// @Summary 更新CloudPost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.CloudPost true "更新CloudPost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cloudPost/updateCloudPost [put]
func (cloudPostApi *CloudPostApi) UpdateCloudPost(c *gin.Context) {
	var cloudPost cloud.CloudPost
	err := c.ShouldBindJSON(&cloudPost)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cloudPost.UpdatedBy = utils.GetUserID(c)
	if err := cloudPostService.UpdateCloudPost(cloudPost); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCloudPost 用id查询CloudPost
// @Tags CloudPost
// @Summary 用id查询CloudPost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloud.CloudPost true "用id查询CloudPost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cloudPost/findCloudPost [get]
func (cloudPostApi *CloudPostApi) FindCloudPost(c *gin.Context) {
	var cloudPost cloud.CloudPost
	err := c.ShouldBindQuery(&cloudPost)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recloudPost, err := cloudPostService.GetCloudPost(cloudPost.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recloudPost": recloudPost}, c)
	}
}

// GetCloudPostList 分页获取CloudPost列表
// @Tags CloudPost
// @Summary 分页获取CloudPost列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloudReq.CloudPostSearch true "分页获取CloudPost列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cloudPost/getCloudPostList [get]
func (cloudPostApi *CloudPostApi) GetCloudPostList(c *gin.Context) {
	var pageInfo cloudReq.CloudPostSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cloudPostService.GetCloudPostInfoList(pageInfo); err != nil {
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
