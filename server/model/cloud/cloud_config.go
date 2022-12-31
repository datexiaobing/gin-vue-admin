// 自动生成模板CloudConfig
package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CloudConfig 结构体
type CloudConfig struct {
	global.GVA_MODEL
	TransDrawtextColor    string `json:"transDrawtextColor" form:"transDrawtextColor" gorm:"column:trans_drawtext_color;comment:跑马灯颜色;size:255;"`
	TransDrawtextDuration *int   `json:"transDrawtextDuration" form:"transDrawtextDuration" gorm:"column:trans_drawtext_duration;comment:滚动时长;size:10;"`
	TransDrawtextFontsize *int   `json:"transDrawtextFontsize" form:"transDrawtextFontsize" gorm:"column:trans_drawtext_fontsize;comment:跑马灯大小;size:10;"`
	TransDrawtextInterval *int   `json:"transDrawtextInterval" form:"transDrawtextInterval" gorm:"column:trans_drawtext_interval;comment:滚动间隔;size:10;"`
	TransDrawtextPosition *int   `json:"transDrawtextPosition" form:"transDrawtextPosition" gorm:"column:trans_drawtext_position;comment:跑马灯位置;size:10;"`
	TransDrawtextString   string `json:"transDrawtextString" form:"transDrawtextString" gorm:"column:trans_drawtext_string;comment:滚动文字;size:255;"`
	TransResolution       *int   `json:"transResolution" form:"transResolution" gorm:"column:trans_resolution;comment:分辨率;size:10;"`
	AliEndPoint           string `json:"aliEndPoint" form:"aliEndPoint" gorm:"column:ali_end_point;comment:存储区域;size:255;"`
	AliAccessKeyId        string `json:"aliAccessKeyId" form:"aliAccessKeyId" gorm:"column:ali_access_key_id;comment:accessKeyId;size:255;"`
	AliAccessKeySecret    string `json:"aliAccessKeySecret" form:"aliAccessKeySecret" gorm:"column:ali_access_key_secret;comment:aliAccessKeySecret;size:255;"`
	AliBucketName         string `json:"aliBucketName" form:"aliBucketName" gorm:"column:ali_bucket_name;comment:aliBucketName;size:255;"`
	TransSeektimeHeard    *int   `json:"transSeektimeHeard" form:"transSeektimeHeard" gorm:"column:trans_seektime_heard;comment:跳过片头;size:10;"`
	TransSeektimeTail     *int   `json:"transSeektimeTail" form:"transSeektimeTail" gorm:"column:trans_seektime_tail;comment:跳过片尾;size:10;"`
	OssStatus             string `json:"ossStatus" form:"ossStatus" gorm:"column:oss_status;comment:1七牛2阿里;default:1"`
	CreatedBy             uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy             uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy             uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName CloudConfig 表名
func (CloudConfig) TableName() string {
	return "cloud_config"
}
