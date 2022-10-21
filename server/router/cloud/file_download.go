package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileDownloadRouter struct {
}

// InitFileDownloadRouter 初始化 FileDownload 路由信息
func (s *FileDownloadRouter) InitFileDownloadRouter(Router *gin.RouterGroup) {
	fileDownloadRouter := Router.Group("fileDownload").Use(middleware.OperationRecord())
	fileDownloadRouterWithoutRecord := Router.Group("fileDownload")
	var fileDownloadApi = v1.ApiGroupApp.CloudApiGroup.FileDownloadApi
	{
		fileDownloadRouter.POST("createFileDownload", fileDownloadApi.CreateFileDownload)   // 新建FileDownload
		fileDownloadRouter.DELETE("deleteFileDownload", fileDownloadApi.DeleteFileDownload) // 删除FileDownload
		fileDownloadRouter.DELETE("deleteFileDownloadByIds", fileDownloadApi.DeleteFileDownloadByIds) // 批量删除FileDownload
		fileDownloadRouter.PUT("updateFileDownload", fileDownloadApi.UpdateFileDownload)    // 更新FileDownload
	}
	{
		fileDownloadRouterWithoutRecord.GET("findFileDownload", fileDownloadApi.FindFileDownload)        // 根据ID获取FileDownload
		fileDownloadRouterWithoutRecord.GET("getFileDownloadList", fileDownloadApi.GetFileDownloadList)  // 获取FileDownload列表
	}
}
