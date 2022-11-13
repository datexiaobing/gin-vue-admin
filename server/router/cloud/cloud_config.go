package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CloudConfigRouter struct {
}

// InitCloudConfigRouter 初始化 CloudConfig 路由信息
func (s *CloudConfigRouter) InitCloudConfigRouter(Router *gin.RouterGroup) {
	cloudConfigRouter := Router.Group("cloudConfig").Use(middleware.OperationRecord())
	cloudConfigRouterWithoutRecord := Router.Group("cloudConfig")
	var cloudConfigApi = v1.ApiGroupApp.CloudApiGroup.CloudConfigApi
	{
		cloudConfigRouter.POST("createCloudConfig", cloudConfigApi.CreateCloudConfig)   // 新建CloudConfig
		cloudConfigRouter.DELETE("deleteCloudConfig", cloudConfigApi.DeleteCloudConfig) // 删除CloudConfig
		cloudConfigRouter.DELETE("deleteCloudConfigByIds", cloudConfigApi.DeleteCloudConfigByIds) // 批量删除CloudConfig
		cloudConfigRouter.PUT("updateCloudConfig", cloudConfigApi.UpdateCloudConfig)    // 更新CloudConfig
	}
	{
		cloudConfigRouterWithoutRecord.GET("findCloudConfig", cloudConfigApi.FindCloudConfig)        // 根据ID获取CloudConfig
		cloudConfigRouterWithoutRecord.GET("getCloudConfigList", cloudConfigApi.GetCloudConfigList)  // 获取CloudConfig列表
	}
}
