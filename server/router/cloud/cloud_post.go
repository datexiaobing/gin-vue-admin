package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CloudPostRouter struct {
}

// InitCloudPostRouter 初始化 CloudPost 路由信息
func (s *CloudPostRouter) InitCloudPostRouter(Router *gin.RouterGroup) {
	cloudPostRouter := Router.Group("cloudPost").Use(middleware.OperationRecord())
	cloudPostRouterWithoutRecord := Router.Group("cloudPost")
	var cloudPostApi = v1.ApiGroupApp.CloudApiGroup.CloudPostApi
	{
		cloudPostRouter.POST("createCloudPost", cloudPostApi.CreateCloudPost)   // 新建CloudPost
		cloudPostRouter.DELETE("deleteCloudPost", cloudPostApi.DeleteCloudPost) // 删除CloudPost
		cloudPostRouter.DELETE("deleteCloudPostByIds", cloudPostApi.DeleteCloudPostByIds) // 批量删除CloudPost
		cloudPostRouter.PUT("updateCloudPost", cloudPostApi.UpdateCloudPost)    // 更新CloudPost
	}
	{
		cloudPostRouterWithoutRecord.GET("findCloudPost", cloudPostApi.FindCloudPost)        // 根据ID获取CloudPost
		cloudPostRouterWithoutRecord.GET("getCloudPostList", cloudPostApi.GetCloudPostList)  // 获取CloudPost列表
	}
}
