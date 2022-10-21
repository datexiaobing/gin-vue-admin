// 自动生成模板VideoCategory
package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// VideoCategory 结构体
type VideoCategory struct {
      global.GVA_MODEL
      CategoryName  string `json:"categoryName" form:"categoryName" gorm:"column:category_name;comment:分类名;size:255;"`
      CategoryPic  string `json:"categoryPic" form:"categoryPic" gorm:"column:category_pic;comment:分类图片;size:255;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName VideoCategory 表名
func (VideoCategory) TableName() string {
  return "video_category"
}

