// 自动生成模板FileDownload
package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// FileDownload 结构体
type FileDownload struct {
      global.GVA_MODEL
      FileDegree  *int `json:"fileDegree" form:"fileDegree" gorm:"column:file_degree;comment:0根目录1一层目录2二层目录;size:10;"`
      FileIsDir  *int `json:"fileIsDir" form:"fileIsDir" gorm:"column:file_is_dir;comment:1dir0文件;size:10;"`
      FileLastDir  string `json:"fileLastDir" form:"fileLastDir" gorm:"column:file_last_dir;comment:最后一层dir名;size:255;"`
      FileName  string `json:"fileName" form:"fileName" gorm:"column:file_name;comment:文件（夹）名;size:255;"`
      FilePath  string `json:"filePath" form:"filePath" gorm:"column:file_path;comment:文件路径;size:255;"`
      FileSize  *float64 `json:"fileSize" form:"fileSize" gorm:"column:file_size;comment:文件大小;"`
      FileType  string `json:"fileType" form:"fileType" gorm:"column:file_type;comment:1video2other;size:255;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName FileDownload 表名
func (FileDownload) TableName() string {
  return "file_download"
}

