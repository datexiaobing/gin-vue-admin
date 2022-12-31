package cloud

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VideoListApi struct {
}

var videoListService = service.ServiceGroupApp.CloudServiceGroup.VideoListService

// CreateVideoList 创建VideoList
// @Tags VideoList
// @Summary 创建VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoList true "创建VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoList/createVideoList [post]
func (videoListApi *VideoListApi) CreateVideoList(c *gin.Context) {
	var videoList cloud.VideoList
	err := c.ShouldBindJSON(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	videoList.CreatedBy = utils.GetUserID(c)
	if err := videoListService.CreateVideoList(videoList); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteVideoList 删除VideoList
// @Tags VideoList
// @Summary 删除VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoList true "删除VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoList/deleteVideoList [delete]
func (videoListApi *VideoListApi) DeleteVideoList(c *gin.Context) {
	var videoList cloud.VideoList
	err := c.ShouldBindJSON(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	videoList.DeletedBy = utils.GetUserID(c)
	if err := videoListService.DeleteVideoList(videoList); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteVideoListByIds 批量删除VideoList
// @Tags VideoList
// @Summary 批量删除VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /videoList/deleteVideoListByIds [delete]
func (videoListApi *VideoListApi) DeleteVideoListByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := videoListService.DeleteVideoListByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateVideoList 更新VideoList
// @Tags VideoList
// @Summary 更新VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cloud.VideoList true "更新VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoList/updateVideoList [put]
func (videoListApi *VideoListApi) UpdateVideoList(c *gin.Context) {
	var videoList cloud.VideoList
	err := c.ShouldBindJSON(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	videoList.UpdatedBy = utils.GetUserID(c)
	if err := videoListService.UpdateVideoList(videoList); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindVideoList 用id查询VideoList
// @Tags VideoList
// @Summary 用id查询VideoList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloud.VideoList true "用id查询VideoList"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoList/findVideoList [get]
func (videoListApi *VideoListApi) FindVideoList(c *gin.Context) {
	var videoList cloud.VideoList
	err := c.ShouldBindQuery(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if revideoList, err := videoListService.GetVideoList(videoList.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"revideoList": revideoList}, c)
	}
}

// GetVideoListList 分页获取VideoList列表
// @Tags VideoList
// @Summary 分页获取VideoList列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cloudReq.VideoListSearch true "分页获取VideoList列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoList/getVideoListList [get]
func (videoListApi *VideoListApi) GetVideoListList(c *gin.Context) {
	var pageInfo cloudReq.VideoListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := videoListService.GetVideoListInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// active search
func (videoListApi *VideoListApi) GetVideoListListActive(c *gin.Context) {
	var pageInfo cloudReq.VideoListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := videoListService.GetVideoListInfoListActive(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// getVideosStatus
func (videoListApi *VideoListApi) GetVideosStatus(c *gin.Context) {
	var pageInfo cloudReq.VideoListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := videoListService.GetVideosStatus(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// waiting search
func (videoListApi *VideoListApi) GetVideoListListWaiting(c *gin.Context) {
	var pageInfo cloudReq.VideoListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := videoListService.GetVideoListInfoListWaiting(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// stop search
func (videoListApi *VideoListApi) GetVideoListListStop(c *gin.Context) {
	var pageInfo cloudReq.VideoListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := videoListService.GetVideoListInfoListStop(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// pause
func (videoListApi *VideoListApi) GetVideoListListPause(c *gin.Context) {
	var videoList request.GridsReq
	err := c.ShouldBindJSON(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := videoListService.GetVideoListInfoListPause(videoList); err != nil {
		global.GVA_LOG.Error("暂停任务失败!", zap.Error(err))
		response.FailWithMessage("暂停任务失败，请稍后再试", c)
	} else {
		response.OkWithMessage("暂停任务成功", c)
	}
}

// unpause
func (videoListApi *VideoListApi) GetVideoListListUnpause(c *gin.Context) {
	var videoList request.GridsReq
	err := c.ShouldBindJSON(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := videoListService.GetVideoListInfoListUnpause(videoList); err != nil {
		global.GVA_LOG.Error("恢复任务失败!", zap.Error(err))
		response.FailWithMessage("恢复任务失败", c)
	} else {
		response.OkWithMessage("恢复任务成功", c)
	}
}

// remove
func (videoListApi *VideoListApi) GetVideoListListRemove(c *gin.Context) {
	var videoList request.GridsReq
	err := c.ShouldBindJSON(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := videoListService.GetVideoListInfoListRemove(videoList); err != nil {
		global.GVA_LOG.Error("删除任务失败!", zap.Error(err))
		response.FailWithMessage("删除任务失败,请稍后再试", c)
	} else {
		response.OkWithMessage("删除任务成功", c)
	}
}

// upload 种子
func (videoListApi *VideoListApi) GetVideoListListUpload(c *gin.Context) {
	// fmt.Println("+++++++++++++++")
	file, _ := c.FormFile("file") // 获取文件
	// fmt.Println(file.Filename)

	unix_int := time.Now().Unix()                    // 时间戳，int类型
	time_unix_str := strconv.FormatInt(unix_int, 10) // 讲int类型转为string类型，方便拼接，使用sprinf也可以
	// 防止重名文件
	// base_path := "D:\\myWork\\korea\\cloud\\gin-vue-admin\\server\\upload\\"
	base_path := "/home/cloud/upload/"
	sysType := runtime.GOOS
	if sysType == "windows" {
		fmt.Println("Windows system")
		base_path = "D:\\myWork\\korea\\cloud\\gin-vue-admin\\server\\upload\\"
	}
	file_path := base_path + time_unix_str + file.Filename // 设置保存文件的路径，不要忘了后面的文件名

	c.SaveUploadedFile(file, file_path) // 保存文件
	// fmt.Println("file_path:", file_path)

	videoListService.DownLoadTorrent(file_path)

	c.String(http.StatusOK, "上传成功")

}

// file search
func (videoListApi *VideoListApi) GetVideoListListFile(c *gin.Context) {
	var pageInfo cloudReq.VideoListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := videoListService.GetVideoListInfoListFile(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// file rename
func (videoListApi *VideoListApi) RenameFiles(c *gin.Context) {
	var videoList request.FileReq
	err := c.ShouldBindJSON(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := videoListService.RenameFile(videoList.DownloadPath, videoList.FileName); err != nil {
		global.GVA_LOG.Error("移动文件失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("移动文件成功", c)
	}
}

// file search
func (videoListApi *VideoListApi) GetVideoListListFileDone(c *gin.Context) {
	var pageInfo cloudReq.VideoListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := videoListService.GetVideoListInfoListFileDone(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// file rename
func (videoListApi *VideoListApi) ChangeFileName(c *gin.Context) {
	var videoList request.FileReq
	err := c.ShouldBindJSON(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := videoListService.ChangeFileName(videoList.DownloadPath, videoList.FileName); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更名成功", c)
	}
}

// delete file
func (videoListApi *VideoListApi) DeleteFile(c *gin.Context) {
	var videoList request.FileReq
	err := c.ShouldBindJSON(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := videoListService.DeleteFile(videoList.DownloadPath, videoList.FileName); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("更名成功", c)
	}
}

// down m3u8
func (videoListApi *VideoListApi) DownM3u8(c *gin.Context) {
	var videoList request.FileReq
	err := c.ShouldBindJSON(&videoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := videoListService.DownM3u8(videoList.DownloadPath, videoList.FileName); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("正在下载...", c)
	}
}
