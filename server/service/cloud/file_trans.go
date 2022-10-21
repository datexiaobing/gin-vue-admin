package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
    "gorm.io/gorm"
)

type FileTransService struct {
}

// CreateFileTrans 创建FileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileTransService *FileTransService) CreateFileTrans(fileTrans cloud.FileTrans) (err error) {
	err = global.GVA_DB.Create(&fileTrans).Error
	return err
}

// DeleteFileTrans 删除FileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileTransService *FileTransService)DeleteFileTrans(fileTrans cloud.FileTrans) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cloud.FileTrans{}).Where("id = ?", fileTrans.ID).Update("deleted_by", fileTrans.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&fileTrans).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteFileTransByIds 批量删除FileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileTransService *FileTransService)DeleteFileTransByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&cloud.FileTrans{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&cloud.FileTrans{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateFileTrans 更新FileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileTransService *FileTransService)UpdateFileTrans(fileTrans cloud.FileTrans) (err error) {
	err = global.GVA_DB.Save(&fileTrans).Error
	return err
}

// GetFileTrans 根据id获取FileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileTransService *FileTransService)GetFileTrans(id uint) (fileTrans cloud.FileTrans, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fileTrans).Error
	return
}

// GetFileTransInfoList 分页获取FileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileTransService *FileTransService)GetFileTransInfoList(info cloudReq.FileTransSearch) (list []cloud.FileTrans, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cloud.FileTrans{})
    var fileTranss []cloud.FileTrans
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.TransOutName != "" {
        db = db.Where("trans_out_name LIKE ?","%"+ info.TransOutName+"%")
    }
    if info.TransStatus != nil {
        db = db.Where("trans_status = ?",info.TransStatus)
    }
    if info.TransType != nil {
        db = db.Where("trans_type = ?",info.TransType)
    }
    if info.TransOssStatus != nil {
        db = db.Where("trans_oss_status = ?",info.TransOssStatus)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&fileTranss).Error
	return  fileTranss, total, err
}
