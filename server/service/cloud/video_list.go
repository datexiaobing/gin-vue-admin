package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
    "gorm.io/gorm"
)

type VideoListService struct {
}

// CreateVideoList 创建VideoList记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoListService *VideoListService) CreateVideoList(videoList cloud.VideoList) (err error) {
	err = global.GVA_DB.Create(&videoList).Error
	return err
}

// DeleteVideoList 删除VideoList记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoListService *VideoListService)DeleteVideoList(videoList cloud.VideoList) (err error) {
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
func (videoListService *VideoListService)DeleteVideoListByIds(ids request.IdsReq,deleted_by uint) (err error) {
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
func (videoListService *VideoListService)UpdateVideoList(videoList cloud.VideoList) (err error) {
	err = global.GVA_DB.Save(&videoList).Error
	return err
}

// GetVideoList 根据id获取VideoList记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoListService *VideoListService)GetVideoList(id uint) (videoList cloud.VideoList, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&videoList).Error
	return
}

// GetVideoListInfoList 分页获取VideoList记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoListService *VideoListService)GetVideoListInfoList(info cloudReq.VideoListSearch) (list []cloud.VideoList, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cloud.VideoList{})
    var videoLists []cloud.VideoList
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.VideoDownloadStatus != nil {
        db = db.Where("video_download_status = ?",info.VideoDownloadStatus)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&videoLists).Error
	return  videoLists, total, err
}
