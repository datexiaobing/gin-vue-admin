package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
    "gorm.io/gorm"
)

type FileInputService struct {
}

// CreateFileInput 创建FileInput记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileInputService *FileInputService) CreateFileInput(fileInput cloud.FileInput) (err error) {
	err = global.GVA_DB.Create(&fileInput).Error
	return err
}

// DeleteFileInput 删除FileInput记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileInputService *FileInputService)DeleteFileInput(fileInput cloud.FileInput) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cloud.FileInput{}).Where("id = ?", fileInput.ID).Update("deleted_by", fileInput.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&fileInput).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteFileInputByIds 批量删除FileInput记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileInputService *FileInputService)DeleteFileInputByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cloud.FileInput{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&cloud.FileInput{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateFileInput 更新FileInput记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileInputService *FileInputService)UpdateFileInput(fileInput cloud.FileInput) (err error) {
	err = global.GVA_DB.Save(&fileInput).Error
	return err
}

// GetFileInput 根据id获取FileInput记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileInputService *FileInputService)GetFileInput(id uint) (fileInput cloud.FileInput, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fileInput).Error
	return
}

// GetFileInputInfoList 分页获取FileInput记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileInputService *FileInputService)GetFileInputInfoList(info cloudReq.FileInputSearch) (list []cloud.FileInput, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cloud.FileInput{})
    var fileInputs []cloud.FileInput
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.FileName != "" {
        db = db.Where("file_name LIKE ?","%"+ info.FileName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&fileInputs).Error
	return  fileInputs, total, err
}
