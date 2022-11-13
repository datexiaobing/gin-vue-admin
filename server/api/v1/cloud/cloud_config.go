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

type CloudConfigApi struct {
}

var cloudConfigService = service.ServiceGroupApp.CloudServiceGroup.CloudConfigService


// CreateCloudConfig 创建CloudConfig
// @Tags CloudConfig
// @Summary 创建CloudConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.CloudConfig true "创建CloudConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cloudConfig/createCloudConfig [post]
func (cloudConfigApi *CloudConfigApi) CreateCloudConfig(c *gin.Context) {
	var cloudConfig cloud.CloudConfig
	err := c.ShouldBindJSON(&cloudConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    cloudConfig.CreatedBy = utils.GetUserID(c)
	if err := cloudConfigService.CreateCloudConfig(cloudConfig); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCloudConfig 删除CloudConfig
// @Tags CloudConfig
// @Summary 删除CloudConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.CloudConfig true "删除CloudConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cloudConfig/deleteCloudConfig [delete]
func (cloudConfigApi *CloudConfigApi) DeleteCloudConfig(c *gin.Context) {
	var cloudConfig cloud.CloudConfig
	err := c.ShouldBindJSON(&cloudConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    cloudConfig.DeletedBy = utils.GetUserID(c)
	if err := cloudConfigService.DeleteCloudConfig(cloudConfig); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCloudConfigByIds 批量删除CloudConfig
// @Tags CloudConfig
// @Summary 批量删除CloudConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CloudConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cloudConfig/deleteCloudConfigByIds [delete]
func (cloudConfigApi *CloudConfigApi) DeleteCloudConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := cloudConfigService.DeleteCloudConfigByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCloudConfig 更新CloudConfig
// @Tags CloudConfig
// @Summary 更新CloudConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.CloudConfig true "更新CloudConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cloudConfig/updateCloudConfig [put]
func (cloudConfigApi *CloudConfigApi) UpdateCloudConfig(c *gin.Context) {
	var cloudConfig cloud.CloudConfig
	err := c.ShouldBindJSON(&cloudConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    cloudConfig.UpdatedBy = utils.GetUserID(c)
	if err := cloudConfigService.UpdateCloudConfig(cloudConfig); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCloudConfig 用id查询CloudConfig
// @Tags CloudConfig
// @Summary 用id查询CloudConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloud.CloudConfig true "用id查询CloudConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cloudConfig/findCloudConfig [get]
func (cloudConfigApi *CloudConfigApi) FindCloudConfig(c *gin.Context) {
	var cloudConfig cloud.CloudConfig
	err := c.ShouldBindQuery(&cloudConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recloudConfig, err := cloudConfigService.GetCloudConfig(cloudConfig.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recloudConfig": recloudConfig}, c)
	}
}

// GetCloudConfigList 分页获取CloudConfig列表
// @Tags CloudConfig
// @Summary 分页获取CloudConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloudReq.CloudConfigSearch true "分页获取CloudConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cloudConfig/getCloudConfigList [get]
func (cloudConfigApi *CloudConfigApi) GetCloudConfigList(c *gin.Context) {
	var pageInfo cloudReq.CloudConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cloudConfigService.GetCloudConfigInfoList(pageInfo); err != nil {
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
