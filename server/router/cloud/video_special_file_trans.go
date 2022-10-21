package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VideoSpecialFileTransRouter struct {
}

// InitVideoSpecialFileTransRouter 初始化 VideoSpecialFileTrans 路由信息
func (s *VideoSpecialFileTransRouter) InitVideoSpecialFileTransRouter(Router *gin.RouterGroup) {
	videoSpecialFileTransRouter := Router.Group("videoSpecialFileTrans").Use(middleware.OperationRecord())
	videoSpecialFileTransRouterWithoutRecord := Router.Group("videoSpecialFileTrans")
	var videoSpecialFileTransApi = v1.ApiGroupApp.CloudApiGroup.VideoSpecialFileTransApi
	{
		videoSpecialFileTransRouter.POST("createVideoSpecialFileTrans", videoSpecialFileTransApi.CreateVideoSpecialFileTrans)   // 新建VideoSpecialFileTrans
		videoSpecialFileTransRouter.DELETE("deleteVideoSpecialFileTrans", videoSpecialFileTransApi.DeleteVideoSpecialFileTrans) // 删除VideoSpecialFileTrans
		videoSpecialFileTransRouter.DELETE("deleteVideoSpecialFileTransByIds", videoSpecialFileTransApi.DeleteVideoSpecialFileTransByIds) // 批量删除VideoSpecialFileTrans
		videoSpecialFileTransRouter.PUT("updateVideoSpecialFileTrans", videoSpecialFileTransApi.UpdateVideoSpecialFileTrans)    // 更新VideoSpecialFileTrans
	}
	{
		videoSpecialFileTransRouterWithoutRecord.GET("findVideoSpecialFileTrans", videoSpecialFileTransApi.FindVideoSpecialFileTrans)        // 根据ID获取VideoSpecialFileTrans
		videoSpecialFileTransRouterWithoutRecord.GET("getVideoSpecialFileTransList", videoSpecialFileTransApi.GetVideoSpecialFileTransList)  // 获取VideoSpecialFileTrans列表
	}
}
