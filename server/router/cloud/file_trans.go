package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileTransRouter struct {
}

// InitFileTransRouter 初始化 FileTrans 路由信息
func (s *FileTransRouter) InitFileTransRouter(Router *gin.RouterGroup) {
	fileTransRouter := Router.Group("fileTrans").Use(middleware.OperationRecord())
	fileTransRouterWithoutRecord := Router.Group("fileTrans")
	var fileTransApi = v1.ApiGroupApp.CloudApiGroup.FileTransApi
	{
		fileTransRouter.POST("createFileTrans", fileTransApi.CreateFileTrans)   // 新建FileTrans
		fileTransRouter.DELETE("deleteFileTrans", fileTransApi.DeleteFileTrans) // 删除FileTrans
		fileTransRouter.DELETE("deleteFileTransByIds", fileTransApi.DeleteFileTransByIds) // 批量删除FileTrans
		fileTransRouter.PUT("updateFileTrans", fileTransApi.UpdateFileTrans)    // 更新FileTrans
	}
	{
		fileTransRouterWithoutRecord.GET("findFileTrans", fileTransApi.FindFileTrans)        // 根据ID获取FileTrans
		fileTransRouterWithoutRecord.GET("getFileTransList", fileTransApi.GetFileTransList)  // 获取FileTrans列表
	}
}
