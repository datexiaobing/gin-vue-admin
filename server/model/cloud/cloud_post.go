// 自动生成模板CloudPost
package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CloudPost 结构体
type CloudPost struct {
	global.GVA_MODEL
	PostNum    int  `json:"postNum" form:"postNum" gorm:"column:post_num;comment:总推送视频数;"`
	SuccessNum int  `json:"successNum" form:"successNum" gorm:"column:success_num;comment:推送成功的视频数;"`
	CreatedBy  uint `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint `gorm:"column:deleted_by;comment:删除者"`
}

// TableName CloudPost 表名
func (CloudPost) TableName() string {
	return "cloud_post"
}
