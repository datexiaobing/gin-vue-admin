package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type VideoSpecialFileTransService struct {
}

// CreateVideoSpecialFileTrans 创建VideoSpecialFileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoSpecialFileTransService *VideoSpecialFileTransService) CreateVideoSpecialFileTrans(videoSpecialFileTrans cloud.VideoSpecialFileTrans) (err error) {
	err = global.GVA_DB.Create(&videoSpecialFileTrans).Error
	return err
}

// DeleteVideoSpecialFileTrans 删除VideoSpecialFileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoSpecialFileTransService *VideoSpecialFileTransService) DeleteVideoSpecialFileTrans(videoSpecialFileTrans cloud.VideoSpecialFileTrans) (err error) {
	err = global.GVA_DB.Delete(&videoSpecialFileTrans).Error
	return err
}

// DeleteVideoSpecialFileTransByIds 批量删除VideoSpecialFileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoSpecialFileTransService *VideoSpecialFileTransService) DeleteVideoSpecialFileTransByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Delete(&[]cloud.VideoSpecialFileTrans{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateVideoSpecialFileTrans 更新VideoSpecialFileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoSpecialFileTransService *VideoSpecialFileTransService) UpdateVideoSpecialFileTrans(videoSpecialFileTrans cloud.VideoSpecialFileTrans) (err error) {
	err = global.GVA_DB.Save(&videoSpecialFileTrans).Error
	return err
}

// GetVideoSpecialFileTrans 根据id获取VideoSpecialFileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoSpecialFileTransService *VideoSpecialFileTransService) GetVideoSpecialFileTrans(id uint) (videoSpecialFileTrans cloud.VideoSpecialFileTrans, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&videoSpecialFileTrans).Error
	return
}

// GetVideoSpecialFileTransInfoList 分页获取VideoSpecialFileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (videoSpecialFileTransService *VideoSpecialFileTransService) GetVideoSpecialFileTransInfoList(info cloudReq.VideoSpecialFileTransSearch) (list []cloud.VideoSpecialFileTrans, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cloud.VideoSpecialFileTrans{})
	var videoSpecialFileTranss []cloud.VideoSpecialFileTrans
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.VideoSpecialId != 0 {
		db = db.Where("video_special_id = ?", info.VideoSpecialId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&videoSpecialFileTranss).Error
	return videoSpecialFileTranss, total, err
}
