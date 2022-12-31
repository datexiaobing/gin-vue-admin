package cloud

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	a "github.com/flipped-aurora/gin-vue-admin/server/aria2"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/m3u8d"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type VideoListService struct {
}

// CreateVideoList 创建VideoList记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoListService *VideoListService) CreateVideoList(videoList cloud.VideoList) (err error) {

	downPath := videoList.VideoDownloadPath
	// fmt.Println("down path:", downPath)
	urlSlice := utils.SplitUrls(downPath)
	// headerSlice := utils.SplitUrls(videoList.VideoUrl)
	// fmt.Println(headerSlice, "header", videoList.VideoUrl)
	option := a.Option{}
	// h := headerSlice[0]
	h := videoList.VideoUrl
	option.Header = h //目前只取一个header
	// option.UserAgent = "Chrome"
	for _, v := range urlSlice {
		fmt.Println("start downloading::", v)
		utils.DownloadByUrl(v, &option)
	}

	return nil
	// err = global.GVA_DB.Create(&videoList).Error
	// return err
}

// DeleteVideoList 删除VideoList记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoListService *VideoListService) DeleteVideoList(videoList cloud.VideoList) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cloud.VideoList{}).Where("id = ?", videoList.ID).Update("deleted_by", videoList.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&videoList).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteVideoListByIds 批量删除VideoList记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoListService *VideoListService) DeleteVideoListByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cloud.VideoList{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&cloud.VideoList{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateVideoList 更新VideoList记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoListService *VideoListService) UpdateVideoList(videoList cloud.VideoList) (err error) {
	err = global.GVA_DB.Save(&videoList).Error
	return err
}

// GetVideoList 根据id获取VideoList记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoListService *VideoListService) GetVideoList(id uint) (videoList cloud.VideoList, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&videoList).Error
	return
}

// GetVideoListInfoList 分页获取VideoList记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoListService *VideoListService) GetVideoListInfoList(info cloudReq.VideoListSearch) (list []cloud.VideoList, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cloud.VideoList{})
	var videoLists []cloud.VideoList
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.VideoDownloadStatus != nil {
		db = db.Where("video_download_status = ?", info.VideoDownloadStatus)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&videoLists).Error
	return videoLists, total, err
}

// active 查询
func (videoListService *VideoListService) GetVideoListInfoListActive(info cloudReq.VideoListSearch) (list interface{}, total int64, err error) {
	// limit := info.PageSize
	// offset := info.PageSize * (info.Page - 1)
	total = 1
	task, err := utils.QueryActiveVideos()
	if len(task) > 0 {
		total = int64(len(task))
	}
	return task, total, err
}

// getVideosStatus  所有状态的数目
func (videoListService *VideoListService) GetVideosStatus(info cloudReq.VideoListSearch) (list interface{}, total int64, err error) {
	// 下载状态查询
	total = 1
	task, _ := utils.GetGlobalStatTask()
	var videoGlobalStat cloud.VideoGlobalStat
	videoGlobalStat.GlobalStatTaskData = *task
	// 转换状态查询
	path := "/home/cloud/mp4"
	files, _ := ioutil.ReadDir(path)
	videoGlobalStat.WaitingNum = len(files) //等于刚移动过来的视频个数
	var fileTrans []cloud.FileTrans
	_ = global.GVA_DB.Model(cloud.FileTrans{}).Where("trans_status = ?", 1).Find(&fileTrans).Error
	videoGlobalStat.ActiviteNum = len(fileTrans)
	// 已完成的视频数量=trans_file表中所有的所有视频数量
	var fileTrans1 []cloud.FileTrans
	err = global.GVA_DB.Model(cloud.FileTrans{}).Where("trans_status = ?", 2).Find(&fileTrans1).Error
	videoGlobalStat.DoneNum = len(fileTrans1)
	return videoGlobalStat, total, err
}

// waiting 查询
func (videoListService *VideoListService) GetVideoListInfoListWaiting(info cloudReq.VideoListSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	total = 1
	task, err := utils.QueryWaitingVideos(offset, limit)
	if len(task) > 0 {
		total = int64(len(task))
	}
	return task, total, err
}

// stop 查询
func (videoListService *VideoListService) GetVideoListInfoListStop(info cloudReq.VideoListSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	total = 1
	task, err := utils.QueryStopVideos(offset, limit)
	if len(task) > 0 {
		total = int64(len(task))
	}
	return task, total, err
}

// 暂停下载
func (videoListService *VideoListService) GetVideoListInfoListPause(grids request.GridsReq) (err error) {
	for _, v := range grids.Grids {
		fmt.Println(v)
		err = utils.PauseVideo(v)
	}
	return err
}

// 恢复下载
func (videoListService *VideoListService) GetVideoListInfoListUnpause(grids request.GridsReq) (err error) {
	for _, v := range grids.Grids {
		fmt.Println(v)
		err = utils.UnpauseVideo(v)
	}
	return err
}

// 删除任务
func (videoListService *VideoListService) GetVideoListInfoListRemove(grids request.GridsReq) (err error) {
	for _, v := range grids.Grids {
		fmt.Println(v)
		err = utils.RemoveTask(v)
	}
	return err
}

// 下载种子
func (videoListService *VideoListService) DownLoadTorrent(path string) (err error) {
	_, err = utils.DownloadTorrent(path)
	return err
}

// file 查询
func (videoListService *VideoListService) GetVideoListInfoListFile(info cloudReq.VideoListSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 下载存放的文件夹
	path := "/home/cloud/aria2/"
	sysType := runtime.GOOS
	if sysType == "windows" {
		fmt.Println("Windows system")
		path = "D:/Program Files (x86)/aria2/aria2-1.34.0/download/"
	}

	if info.VideoUrl == "/" {
		// 主目录
		fmt.Println("主目录")
	} else {
		// 多层目录是否存在问题？
		// path = path + info.VideoUrl
		path = filepath.Join(path, info.VideoUrl)
	}
	total = 1

	files, err := ioutil.ReadDir(path)
	// total = int64(len(files))
	var task []cloud.FileDown
	videoType := []string{"mp4", "3gp", "avi", "flv", "mkv"}
	for _, f := range files {
		// fmt.Println(f)
		temp_task := cloud.FileDown{}
		temp_task.FileName = f.Name()
		temp_task.DownloadPath = filepath.Join(path, f.Name())
		temp_task.FileSize = f.Size()
		temp_task.FileType = "other"
		for _, v := range videoType {
			if strings.Contains(temp_task.FileName, v) {
				temp_task.FileType = "video"
			}
		}
		temp_task.IsDir = f.IsDir()
		temp_task.ModTime = f.ModTime()
		task = append(task, temp_task)
		total += 1
	}
	var top int
	top = offset + limit
	taskout := task
	if len(task) > limit {
		// 文件总数大于pagesize才需要分页
		if len(task) < top {
			top = len(task)
		}
		taskout = task[offset:top]
	}

	return taskout, total, err
}

// move file
func (videoListService *VideoListService) RenameFile(path string, fileName string) (err error) {
	// 移动文件到全部是视频的目录下
	sysType := runtime.GOOS
	des_path := "/home/cloud/mp4/"
	if sysType == "windows" {
		fmt.Println("Windows system")
		des_path = "D:/Program Files (x86)/aria2/aria2-1.34.0/trans/"
	}

	err = os.Rename(path, des_path+fileName)

	return err
}

// file 查询已移动
func (videoListService *VideoListService) GetVideoListInfoListFileDone(info cloudReq.VideoListSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	sysType := runtime.GOOS
	path := "/home/cloud/mp4/"
	if sysType == "windows" {
		fmt.Println("Windows system")
		path = "D:/Program Files (x86)/aria2/aria2-1.34.0/trans/"
	}
	// videoUrl  "/",子文件夹
	// path := "/home/mp4/"
	if info.VideoUrl == "/" {
		// 主目录
		fmt.Println("主目录")
	} else {
		// 多层目录是否存在问题？
		path = filepath.Join(path, info.VideoUrl)
	}
	total = 1

	files, err := ioutil.ReadDir(path)
	// total = int64(len(files))
	var task []cloud.FileDown
	videoType := []string{"mp4", "3gp", "avi", "flv"}
	for _, f := range files {
		// fmt.Println(f)
		temp_task := cloud.FileDown{}
		temp_task.FileName = f.Name()
		temp_task.DownloadPath = filepath.Join(path, f.Name())
		temp_task.FileSize = f.Size()
		temp_task.FileType = "other"
		for _, v := range videoType {
			if strings.Contains(temp_task.FileName, v) {
				temp_task.FileType = "video"
			}
		}
		temp_task.IsDir = f.IsDir()
		temp_task.ModTime = f.ModTime()
		task = append(task, temp_task)
		total += 1
	}
	var top int
	top = offset + limit
	taskout := task
	if len(task) > limit {
		// 文件总数大于pagesize才需要分页
		if len(task) < top {
			top = len(task)
		}
		taskout = task[offset:top]
	}
	return taskout, total, err
}

// change file name
func (videoListService *VideoListService) ChangeFileName(path string, fileName string) (err error) {
	base_path := filepath.Dir(path)
	new_path := filepath.Join(base_path, fileName)
	err = os.Rename(path, new_path)
	return err
}

// delete file
func (videoListService *VideoListService) DeleteFile(path string, fileName string) (err error) {
	err = os.Remove(path)
	return err
}

// down load m3u8
func (videoListService *VideoListService) DownM3u8(path string, fileName string) (err error) {
	var req m3u8d.RunDownload_Req
	req.M3u8Url = path
	savePath := "/home/cloud/mp4/"
	sysType := runtime.GOOS
	if sysType == "windows" {
		fmt.Println("Windows system")
		savePath = "D:/Program Files (x86)/aria2/aria2-1.34.0/download/"
	}

	req.SaveDir = savePath
	req.Insecure = true
	req.ThreadCount = 16
	req.FileName = fileName

	fmt.Println("fileName:", fileName)
	go m3u8d.RunDownload(req)
	return nil
}
