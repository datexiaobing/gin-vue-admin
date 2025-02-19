package cloud

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
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
		videoListRouter.POST("createVideoList", videoListApi.CreateVideoList)             // 新建VideoList
		videoListRouter.DELETE("deleteVideoList", videoListApi.DeleteVideoList)           // 删除VideoList
		videoListRouter.DELETE("deleteVideoListByIds", videoListApi.DeleteVideoListByIds) // 批量删除VideoList
		videoListRouter.PUT("updateVideoList", videoListApi.UpdateVideoList)              // 更新VideoList
		videoListRouter.POST("getVideoListListUnpause", videoListApi.GetVideoListListUnpause)
		videoListRouter.POST("getVideoListListPause", videoListApi.GetVideoListListPause)
		videoListRouter.POST("getVideoListListRemove", videoListApi.GetVideoListListRemove)
		videoListRouter.POST("getVideoListListUpload", videoListApi.GetVideoListListUpload)
		videoListRouter.PUT("renameFile", videoListApi.RenameFiles)         //移动文件
		videoListRouter.POST("changeFileName", videoListApi.ChangeFileName) //移动文件
		videoListRouter.POST("deleteFile", videoListApi.DeleteFile)         //移动文件
		videoListRouter.POST("downM3u8", videoListApi.DownM3u8)

	}
	{
		videoListRouterWithoutRecord.GET("findVideoList", videoListApi.FindVideoList)       // 根据ID获取VideoList
		videoListRouterWithoutRecord.GET("getVideoListList", videoListApi.GetVideoListList) // 获取VideoList列表
		videoListRouterWithoutRecord.GET("getVideoListListActive", videoListApi.GetVideoListListActive)
		videoListRouterWithoutRecord.GET("getVideosStatus", videoListApi.GetVideosStatus) //获取首页下载文件个数
		videoListRouterWithoutRecord.GET("getVideoListListWaiting", videoListApi.GetVideoListListWaiting)
		videoListRouterWithoutRecord.GET("getVideoListListStop", videoListApi.GetVideoListListStop)
		videoListRouterWithoutRecord.GET("getVideoListListFile", videoListApi.GetVideoListListFile)
		videoListRouterWithoutRecord.GET("getVideoListListFileDone", videoListApi.GetVideoListListFileDone)
	}
}
