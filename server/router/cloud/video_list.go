package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VideoListRouter struct {
}

// InitVideoListRouter 初始化 VideoList 路由信息
func (s *VideoListRouter) InitVideoListRouter(Router *gin.RouterGroup) {
	videoListRouter := Router.Group("videoList").Use(middleware.OperationRecord())
	videoListRouterWithoutRecord := Router.Group("videoList")
	var videoListApi = v1.ApiGroupApp.CloudApiGroup.VideoListApi
	{
		videoListRouter.POST("createVideoList", videoListApi.CreateVideoList)   // 新建VideoList
		videoListRouter.DELETE("deleteVideoList", videoListApi.DeleteVideoList) // 删除VideoList
		videoListRouter.DELETE("deleteVideoListByIds", videoListApi.DeleteVideoListByIds) // 批量删除VideoList
		videoListRouter.PUT("updateVideoList", videoListApi.UpdateVideoList)    // 更新VideoList
	}
	{
		videoListRouterWithoutRecord.GET("findVideoList", videoListApi.FindVideoList)        // 根据ID获取VideoList
		videoListRouterWithoutRecord.GET("getVideoListList", videoListApi.GetVideoListList)  // 获取VideoList列表
	}
}
