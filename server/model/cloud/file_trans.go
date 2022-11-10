// 自动生成模板FileTrans
package cloud

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type FileShare struct {
	FileName string `json:"fileName" form:"fileName"`
	M3Url    string `josn:"m3Url" form:"m3Url"`
	PicUrl   string `json:"picUrl" form:"picUrl"`
}

// FileTrans 结构体
type FileTrans struct {
	global.GVA_MODEL
	TransDrawtextSwitch   int     `json:"transDrawtextSwitch" form:"transDrawtextSwitch" gorm:"column:trans_drawtext_switch;comment:是否开启跑马灯;default:0"`
	TransCryp             int     `json:"transCryp" form:"transCryp" gorm:"column:trans_cryp;comment:是否加密;defalut:0;"`
	TransInputName        string  `json:"transInputName" form:"transInputName" gorm:"column:trans_input_name;comment:传入视频名;size:255;"`
	TransOutName          string  `json:"transOutName" form:"transOutName" gorm:"column:trans_out_name;comment:转换后的名;size:255;"`
	TransUuid             string  `json:"transUuid" form:"transUuid" gorm:"column:trans_uuid;comment:uuid;size:255;"`
	TransStatus           int     `json:"transStatus" form:"transStatus" gorm:"column:trans_status;comment:1正在2完成;size:10;default:1"`
	TransType             *int    `json:"transType" form:"transType" gorm:"column:trans_type;comment:专辑;size:10;default:0;"`
	TransTypeNum          *int    `json:"transTypeNum" form:"transTypeNum" gorm:"column:trans_type_num;comment:专辑视频数量;size:10;"`
	TransResolution       int     `json:"transResolution" form:"transResolution" gorm:"column:trans_resolution;comment:分辨率;size:10;default:3"`
	TransDuration         string  `json:"transDuration" form:"transDuration" gorm:"column:trans_duration;comment:时长;size:255;"`
	TransSeektimeHeard    int     `json:"transSeektimeHeard" form:"transSeektimeHeard" gorm:"column:trans_seektime_heard;comment:跳过片头秒;size:10;"`
	TransSeektimeTail     int     `json:"transSeektimeTail" form:"transSeektimeTail" gorm:"column:trans_seektime_tail;comment:跳过片尾;size:10;"`
	TransDrawtextColor    string  `json:"transDrawtextColor" form:"transDrawtextColor" gorm:"column:trans_drawtext_color;comment:跑马灯颜色;size:255;default:white;"`
	TransDrawtextPosition int     `json:"transDrawtextPosition" form:"transDrawtextPosition" gorm:"column:trans_drawtext_position;comment:跑马灯位置1上2下;size:255;"`
	TransDrawtextDuration int     `json:"transDrawtextDuration" form:"transDrawtextDuration" gorm:"column:trans_drawtext_duration;comment:滚动时长;size:255;"`
	TransDrawtextInterval int     `json:"transDrawtextInterval" form:"transDrawtextInterval" gorm:"column:trans_drawtext_interval;comment:滚动间隔;size:255;"`
	TransDrawtextString   string  `json:"transDrawtextString" form:"transDrawtextString" gorm:"column:trans_drawtext_string;comment:跑马灯文字;size:255;"`
	TransDrawtextFontsize int     `json:"transDrawtextFontsize" form:"transDrawtextFontsize" gorm:"column:trans_drawtext_fontsize;comment:文字大小;size:10;"`
	TransSubtitle         string  `json:"transSubtitle" form:"transSubtitle" gorm:"column:trans_subtitle;comment:字幕文件;size:255;"`
	TransSubtitleSwitch   *int    `json:"transSubtitleSwitch" form:"transSubtitleSwitch" gorm:"column:trans_subtitle_switch;comment:字幕开关;size:10;"`
	TransProgressRate     float64 `json:"transProgressRate" form:"transProgressRate" gorm:"column:trans_progress_rate;comment:进度状态;"`
	TransOssStatus        int     `json:"transOssStatus" form:"transOssStatus" gorm:"column:trans_oss_status;comment:0等待上传1已经上传;size:10;"`
	TransOssError         string  `json:"transOssError" form:"transOssError" gorm:"column:trans_oss_error;comment:上传错误;size:255;"`
	CreatedBy             uint    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy             uint    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy             uint    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName FileTrans 表名
func (FileTrans) TableName() string {
	return "file_trans"
}
