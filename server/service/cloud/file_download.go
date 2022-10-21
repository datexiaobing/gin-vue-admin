package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
    "gorm.io/gorm"
)

type FileDownloadService struct {
}

// CreateFileDownload 创建FileDownload记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileDownloadService *FileDownloadService) CreateFileDownload(fileDownload cloud.FileDownload) (err error) {
	err = global.GVA_DB.Create(&fileDownload).Error
	return err
}

// DeleteFileDownload 删除FileDownload记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileDownloadService *FileDownloadService)DeleteFileDownload(fileDownload cloud.FileDownload) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cloud.FileDownload{}).Where("id = ?", fileDownload.ID).Update("deleted_by", fileDownload.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&fileDownload).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteFileDownloadByIds 批量删除FileDownload记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileDownloadService *FileDownloadService)DeleteFileDownloadByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cloud.FileDownload{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&cloud.FileDownload{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateFileDownload 更新FileDownload记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileDownloadService *FileDownloadService)UpdateFileDownload(fileDownload cloud.FileDownload) (err error) {
	err = global.GVA_DB.Save(&fileDownload).Error
	return err
}

// GetFileDownload 根据id获取FileDownload记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileDownloadService *FileDownloadService)GetFileDownload(id uint) (fileDownload cloud.FileDownload, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fileDownload).Error
	return
}

// GetFileDownloadInfoList 分页获取FileDownload记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileDownloadService *FileDownloadService)GetFileDownloadInfoList(info cloudReq.FileDownloadSearch) (list []cloud.FileDownload, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cloud.FileDownload{})
    var fileDownloads []cloud.FileDownload
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.FileDegree != nil {
        db = db.Where("file_degree = ?",info.FileDegree)
    }
    if info.FileName != "" {
        db = db.Where("file_name LIKE ?","%"+ info.FileName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&fileDownloads).Error
	return  fileDownloads, total, err
}
