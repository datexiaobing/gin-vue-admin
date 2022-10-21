package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type VideoSpecialService struct {
}

// CreateVideoSpecial 创建VideoSpecial记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoSpecialService *VideoSpecialService) CreateVideoSpecial(videoSpecial cloud.VideoSpecial) (err error) {
	err = global.GVA_DB.Create(&videoSpecial).Error
	return err
}

// DeleteVideoSpecial 删除VideoSpecial记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoSpecialService *VideoSpecialService) DeleteVideoSpecial(videoSpecial cloud.VideoSpecial) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cloud.VideoSpecial{}).Where("id = ?", videoSpecial.ID).Update("deleted_by", videoSpecial.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&videoSpecial).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteVideoSpecialByIds 批量删除VideoSpecial记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoSpecialService *VideoSpecialService) DeleteVideoSpecialByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cloud.VideoSpecial{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&cloud.VideoSpecial{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateVideoSpecial 更新VideoSpecial记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoSpecialService *VideoSpecialService) UpdateVideoSpecial(videoSpecial cloud.VideoSpecial) (err error) {
	err = global.GVA_DB.Save(&videoSpecial).Error
	return err
}

// GetVideoSpecial 根据id获取VideoSpecial记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoSpecialService *VideoSpecialService) GetVideoSpecial(id uint) (videoSpecial cloud.VideoSpecial, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("SpecialFile").First(&videoSpecial).Error
	return
}

// GetVideoSpecialInfoList 分页获取VideoSpecial记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoSpecialService *VideoSpecialService) GetVideoSpecialInfoList(info cloudReq.VideoSpecialSearch) (list []cloud.VideoSpecial, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cloud.VideoSpecial{})
	var videoSpecials []cloud.VideoSpecial
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.SpecialName != "" {
		db = db.Where("special_name LIKE ?", "%"+info.SpecialName+"%")
	}
	if info.SpecialCategory != nil {
		db = db.Where("special_category = ?", info.SpecialCategory)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&videoSpecials).Error
	return videoSpecials, total, err
}
