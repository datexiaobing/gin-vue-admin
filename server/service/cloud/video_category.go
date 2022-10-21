package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
    "gorm.io/gorm"
)

type VideoCategoryService struct {
}

// CreateVideoCategory 创建VideoCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoCategoryService *VideoCategoryService) CreateVideoCategory(videoCategory cloud.VideoCategory) (err error) {
	err = global.GVA_DB.Create(&videoCategory).Error
	return err
}

// DeleteVideoCategory 删除VideoCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoCategoryService *VideoCategoryService)DeleteVideoCategory(videoCategory cloud.VideoCategory) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cloud.VideoCategory{}).Where("id = ?", videoCategory.ID).Update("deleted_by", videoCategory.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&videoCategory).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteVideoCategoryByIds 批量删除VideoCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoCategoryService *VideoCategoryService)DeleteVideoCategoryByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cloud.VideoCategory{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&cloud.VideoCategory{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateVideoCategory 更新VideoCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoCategoryService *VideoCategoryService)UpdateVideoCategory(videoCategory cloud.VideoCategory) (err error) {
	err = global.GVA_DB.Save(&videoCategory).Error
	return err
}

// GetVideoCategory 根据id获取VideoCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoCategoryService *VideoCategoryService)GetVideoCategory(id uint) (videoCategory cloud.VideoCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&videoCategory).Error
	return
}

// GetVideoCategoryInfoList 分页获取VideoCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoCategoryService *VideoCategoryService)GetVideoCategoryInfoList(info cloudReq.VideoCategorySearch) (list []cloud.VideoCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cloud.VideoCategory{})
    var videoCategorys []cloud.VideoCategory
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.CategoryName != "" {
        db = db.Where("category_name LIKE ?","%"+ info.CategoryName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&videoCategorys).Error
	return  videoCategorys, total, err
}
