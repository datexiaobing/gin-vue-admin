// 自动生成模板VideoSpecial
package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// VideoSpecial 结构体
type VideoSpecial struct {
	global.GVA_MODEL
	SpecialName     string                  `json:"specialName" form:"specialName" gorm:"column:special_name;comment:专辑名称;size:255;"`
	SpecialCategory *int                    `json:"specialCategory" form:"specialCategory" gorm:"column:special_category;comment:分类;size:10;"`
	SpecialVideoNum *int                    `json:"specialVideoNum" form:"specialVideoNum" gorm:"column:special_video_num;comment:视频数量;size:10;"`
	SpecialPic      string                  `json:"specialPic" form:"specialPic" gorm:"column:special_pic;comment:专辑图片;size:255;"`
	CreatedBy       uint                    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy       uint                    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy       uint                    `gorm:"column:deleted_by;comment:删除者"`
	SpecialFile     []VideoSpecialFileTrans `json:"specialFile" form:"specialFile" `
}

// gorm:"many2many:video_special_file_trans;"

// TableName VideoSpecial 表名
func (VideoSpecial) TableName() string {
	return "video_special"
}
