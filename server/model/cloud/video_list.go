// 自动生成模板VideoList
package cloud

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// VideoList 结构体
type VideoList struct {
	global.GVA_MODEL
	VideoGrid           string   `json:"videoGrid" form:"videoGrid" gorm:"column:video_grid;conment:下载id;"`
	VideoTitle          string   `json:"videoTitle" form:"videoTitle" gorm:"column:video_title;comment:名称;size:255;"`
	VideoDownloadPath   string   `json:"videoDownloadPath" form:"videoDownloadPath" gorm:"column:video_download_path;comment:存储地址;size:255;"`
	VideoUrl            string   `json:"videoUrl" form:"videoUrl" gorm:"column:video_url;comment:下载链接;size:255;"`
	VideoPic            string   `json:"videoPic" form:"videoPic" gorm:"column:video_pic;comment:封面;size:255;"`
	VideoDownloadStatus *int     `json:"videoDownloadStatus" form:"videoDownloadStatus" gorm:"column:video_download_status;comment:1未下载2暂停下载3下载完成;size:10;"`
	VideoSize           *float64 `json:"videoSize" form:"videoSize" gorm:"column:video_size;comment:总大小;size:22;"`
	VideoDownloadSize   *float64 `json:"videoDownloadSize" form:"videoDownloadSize" gorm:"column:video_download_size;comment:已下载大小;size:22;"`
	VideoDownloadSpeed  *float64 `json:"videoDownloadSpeed" form:"videoDownloadSpeed" gorm:"column:video_download_speed;comment:下载速率;size:22;"`
	CreatedBy           uint     `gorm:"column:created_by;comment:创建者"`
	UpdatedBy           uint     `gorm:"column:updated_by;comment:更新者"`
	DeletedBy           uint     `gorm:"column:deleted_by;comment:删除者"`
}

type FileDown struct {
	FileName     string    `json:"fileName"`
	IsDir        bool      `json:"isDir"`
	ModTime      time.Time `json:"modTime"`
	DownloadPath string    `json:"downloadPath"`
	FileSize     int64     `json:"fileSize"`
	FileType     string    `json:"fileType"`
}

// TableName VideoList 表名
func (VideoList) TableName() string {
	return "video_list"
}
