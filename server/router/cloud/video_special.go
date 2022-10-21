package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VideoSpecialRouter struct {
}

// InitVideoSpecialRouter 初始化 VideoSpecial 路由信息
func (s *VideoSpecialRouter) InitVideoSpecialRouter(Router *gin.RouterGroup) {
	videoSpecialRouter := Router.Group("videoSpecial").Use(middleware.OperationRecord())
	videoSpecialRouterWithoutRecord := Router.Group("videoSpecial")
	var videoSpecialApi = v1.ApiGroupApp.CloudApiGroup.VideoSpecialApi
	{
		videoSpecialRouter.POST("createVideoSpecial", videoSpecialApi.CreateVideoSpecial)   // 新建VideoSpecial
		videoSpecialRouter.DELETE("deleteVideoSpecial", videoSpecialApi.DeleteVideoSpecial) // 删除VideoSpecial
		videoSpecialRouter.DELETE("deleteVideoSpecialByIds", videoSpecialApi.DeleteVideoSpecialByIds) // 批量删除VideoSpecial
		videoSpecialRouter.PUT("updateVideoSpecial", videoSpecialApi.UpdateVideoSpecial)    // 更新VideoSpecial
	}
	{
		videoSpecialRouterWithoutRecord.GET("findVideoSpecial", videoSpecialApi.FindVideoSpecial)        // 根据ID获取VideoSpecial
		videoSpecialRouterWithoutRecord.GET("getVideoSpecialList", videoSpecialApi.GetVideoSpecialList)  // 获取VideoSpecial列表
	}
}
