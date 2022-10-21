// 自动生成模板VideoSpecialFileTrans
package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// VideoSpecialFileTrans 结构体
type VideoSpecialFileTrans struct {
	global.GVA_MODEL
	VideoSpecialId   uint   `json:"videoSpecialId" form:"videoSpecialId" gorm:"column:video_special_id;comment:;size:10;"`
	FileTransId      uint   `json:"fileTransId" form:"fileTransId" gorm:"column:file_trans_id;comment:;size:10;"`
	FileTransOutName string `json:"fileTransOutName" form:"fileTransOutName" gorm:"column:file_trans_out_name;comment:;size:255;"`
	FileTransUuid    string `json:"fileTransUuid" form:"fileTransUuid" gorm:"column:file_trans_uuid;comment:;size:255;"`
	CreatedBy        uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy        uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy        uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName VideoSpecialFileTrans 表名
func (VideoSpecialFileTrans) TableName() string {
	return "video_special_file_trans"
}
