package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileInputRouter struct {
}

// InitFileInputRouter 初始化 FileInput 路由信息
func (s *FileInputRouter) InitFileInputRouter(Router *gin.RouterGroup) {
	fileInputRouter := Router.Group("fileInput").Use(middleware.OperationRecord())
	fileInputRouterWithoutRecord := Router.Group("fileInput")
	var fileInputApi = v1.ApiGroupApp.CloudApiGroup.FileInputApi
	{
		fileInputRouter.POST("createFileInput", fileInputApi.CreateFileInput)   // 新建FileInput
		fileInputRouter.DELETE("deleteFileInput", fileInputApi.DeleteFileInput) // 删除FileInput
		fileInputRouter.DELETE("deleteFileInputByIds", fileInputApi.DeleteFileInputByIds) // 批量删除FileInput
		fileInputRouter.PUT("updateFileInput", fileInputApi.UpdateFileInput)    // 更新FileInput
	}
	{
		fileInputRouterWithoutRecord.GET("findFileInput", fileInputApi.FindFileInput)        // 根据ID获取FileInput
		fileInputRouterWithoutRecord.GET("getFileInputList", fileInputApi.GetFileInputList)  // 获取FileInput列表
	}
}
