package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VideoCategoryRouter struct {
}

// InitVideoCategoryRouter 初始化 VideoCategory 路由信息
func (s *VideoCategoryRouter) InitVideoCategoryRouter(Router *gin.RouterGroup) {
	videoCategoryRouter := Router.Group("videoCategory").Use(middleware.OperationRecord())
	videoCategoryRouterWithoutRecord := Router.Group("videoCategory")
	var videoCategoryApi = v1.ApiGroupApp.CloudApiGroup.VideoCategoryApi
	{
		videoCategoryRouter.POST("createVideoCategory", videoCategoryApi.CreateVideoCategory)   // 新建VideoCategory
		videoCategoryRouter.DELETE("deleteVideoCategory", videoCategoryApi.DeleteVideoCategory) // 删除VideoCategory
		videoCategoryRouter.DELETE("deleteVideoCategoryByIds", videoCategoryApi.DeleteVideoCategoryByIds) // 批量删除VideoCategory
		videoCategoryRouter.PUT("updateVideoCategory", videoCategoryApi.UpdateVideoCategory)    // 更新VideoCategory
	}
	{
		videoCategoryRouterWithoutRecord.GET("findVideoCategory", videoCategoryApi.FindVideoCategory)        // 根据ID获取VideoCategory
		videoCategoryRouterWithoutRecord.GET("getVideoCategoryList", videoCategoryApi.GetVideoCategoryList)  // 获取VideoCategory列表
	}
}
