package cloud

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type CloudConfigService struct {
}

// CreateCloudConfig 创建CloudConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (cloudConfigService *CloudConfigService) CreateCloudConfig(cloudConfig cloud.CloudConfig) (err error) {
	cloudConfig.ID = 1
	cloudConfig.CreatedAt = time.Now()
	err = global.GVA_DB.Save(&cloudConfig).Error
	return err
}

// DeleteCloudConfig 删除CloudConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (cloudConfigService *CloudConfigService) DeleteCloudConfig(cloudConfig cloud.CloudConfig) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cloud.CloudConfig{}).Where("id = ?", cloudConfig.ID).Update("deleted_by", cloudConfig.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&cloudConfig).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCloudConfigByIds 批量删除CloudConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (cloudConfigService *CloudConfigService) DeleteCloudConfigByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&cloud.CloudConfig{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&cloud.CloudConfig{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCloudConfig 更新CloudConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (cloudConfigService *CloudConfigService) UpdateCloudConfig(cloudConfig cloud.CloudConfig) (err error) {
	err = global.GVA_DB.Save(&cloudConfig).Error
	return err
}

// GetCloudConfig 根据id获取CloudConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (cloudConfigService *CloudConfigService) GetCloudConfig(id uint) (cloudConfig cloud.CloudConfig, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cloudConfig).Error
	return
}

// GetCloudConfigInfoList 分页获取CloudConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (cloudConfigService *CloudConfigService) GetCloudConfigInfoList(info cloudReq.CloudConfigSearch) (list []cloud.CloudConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cloud.CloudConfig{})
	var cloudConfigs []cloud.CloudConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&cloudConfigs).Error
	return cloudConfigs, total, err
}
