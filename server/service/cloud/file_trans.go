package cloud

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/tools"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type FileTransService struct {
}

// 获取转码是否出错
// var messages = make(chan error)

// CreateFileTrans 创建FileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
// 此接口是创建视频转换的接口
func (fileTransService *FileTransService) write(msg float64, file cloud.FileTrans) {
	// 实时更新转换进度，更新数据库
	// fmt.Println("process is:", msg, "file id:", file.TransUuid)
	file.TransProgressRate = msg
	if msg > 99.99 {
		// file.TransStatus = 2
		global.GVA_DB.Table("file_trans").Where("trans_uuid =?", file.TransUuid).Update("trans_status", 2)
	}

	global.GVA_DB.Table("file_trans").Where("trans_uuid=?", file.TransUuid).Update("trans_progress_rate", msg)
}

// 转换视频的函数，多进程开启
func (fileTransService *FileTransService) TransVideo(fileTrans request.TransVideoReq, baseOutPath string) (err error) {
	var tconfig tools.TranscodingConfig
	var wconfig tools.WatermarkConfig
	var vconfig tools.VideoCaptureConfig

	// 定义要转换成的视频 基本配置
	tconfig.OutputType = "m3u8"
	tconfig.InputPath = fileTrans.DownloadPath
	tconfig.OutputPath = filepath.Join(baseOutPath, fileTrans.FileTrans.TransUuid)
	if fileTrans.FileTrans.TransCryp > 0 {
		tconfig.HlsKey = true //暂时设定为不加密，可以设置前端选择
	}
	tconfig.HlsTime = 10        //10s一个ts
	tconfig.ResolutionRatio = 3 //默认分辨率设置为1080p 3:1080P  1:360p 2:720p 4:1440p
	if fileTrans.FileTrans.TransResolution > 0 {
		tconfig.ResolutionRatio = fileTrans.FileTrans.TransResolution
	}
	// 截图设置
	vconfig.Switch = true
	vconfig.StatrTime = "00:01:00" //在1分钟处截图

	// 跑马灯
	if fileTrans.FileTrans.TransDrawtextSwitch > 0 {
		wconfig.Switch = true
	}

	wconfig.Type = 1 //文字
	wconfig.Fontcolor = fileTrans.FileTrans.TransDrawtextColor
	wconfig.Fontsize = fileTrans.FileTrans.TransDrawtextFontsize     //跑马灯大小
	wconfig.RollingSpace = fileTrans.FileTrans.TransDrawtextInterval //间隔10s循环一次
	wconfig.RollingSpeed = 150                                       //滚动时长
	wconfig.RollingStatr = 3
	wconfig.Text = fileTrans.FileTrans.TransDrawtextString
	wconfig.TextType = 1   //滚动
	wconfig.BasicPoint = 3 //2左下 3左上 没用上

	wconfig.XCoordinate = 10 //没用
	wconfig.YCoordinate = 20

	err = tools.Transcoding(tconfig, wconfig, vconfig, fileTrans.FileTrans, fileTransService.write)
	// err 存入messages
	// messages <- err
	if err != nil {
		fmt.Println(err, "转换视频时候故障，退出")
	}
	return err
}

func (fileTransService *FileTransService) CreateFileTrans(fileTrans request.TransVideoReq) (err error) {

	var file cloud.FileTrans
	// linux 环境需要改这里
	baseOutPath := "D:\\myWork\\korea\\cloud\\translate\\transcoding-server\\temp\\testvideo\\out\\a"
	// 获取视频时长，秒
	// ffprobe -v error -show_entries format=duration -of default=noprint_wrappers=1:nokey=1 -i "D:\\Program Files (x86)\\aria2\\aria2-1.34.0\\trans\\1.mp4"
	mediaUrl := fileTrans.DownloadPath
	duration, err := tools.GetMP4Duration(mediaUrl)
	if err != nil {
		fmt.Println(err)
		duration = "00:00:00"
	}
	// fmt.Println(duration)
	fileTrans.FileTrans.TransDuration = duration //视频时长
	fileTrans.FileTrans.TransInputName = fileTrans.FileName
	// output name 由前端上传 不传则和原来一致
	if fileTrans.FileTrans.TransOutName == "" {
		fileTrans.FileTrans.TransOutName = fileTrans.FileName
	}

	uuid := uuid.NewV4()
	fileTrans.FileTrans.TransUuid = uuid.String()

	//开启转换进程 异步
	go fileTransService.TransVideo(fileTrans, baseOutPath)
	// err = <-messages
	// if err != nil {
	// 	fmt.Println(err, "转换视频时候故障，退出")
	// }
	file = fileTrans.FileTrans
	fmt.Println(fileTrans, "success! save data to mysql", file)

	err = global.GVA_DB.Create(&file).Error
	return err
}

// DeleteFileTrans 删除FileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileTransService *FileTransService) DeleteFileTrans(fileTrans cloud.FileTrans) (err error) {
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
func (fileTransService *FileTransService) DeleteFileTransByIds(ids request.IdsReq, deleted_by uint) (err error) {
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
func (fileTransService *FileTransService) UpdateFileTrans(fileTrans cloud.FileTrans) (err error) {
	err = global.GVA_DB.Save(&fileTrans).Error
	return err
}

// GetFileTrans 根据id获取FileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileTransService *FileTransService) GetFileTrans(id uint) (fileTrans cloud.FileTrans, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fileTrans).Error
	return
}

// GetFileTransInfoList 分页获取FileTrans记录
// Author [piexlmax](https://github.com/piexlmax)
func (fileTransService *FileTransService) GetFileTransInfoList(info cloudReq.FileTransSearch) (list []cloud.FileTrans, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&cloud.FileTrans{})
	var fileTranss []cloud.FileTrans
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.TransOutName != "" {
		db = db.Where("trans_out_name LIKE ?", "%"+info.TransOutName+"%")
	}
	if info.TransStatus > 0 {
		db = db.Where("trans_status = ?", info.TransStatus)
	}
	if info.TransType != nil {
		db = db.Where("trans_type = ?", info.TransType)
	}
	if info.TransOssStatus > 0 {
		db = db.Where("trans_oss_status = ?", info.TransOssStatus)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&fileTranss).Error
	return fileTranss, total, err
}

func (fileTransService *FileTransService) GetShareList(share request.IdsShareReq) (list interface{}, err error) {

	// 创建db
	db := global.GVA_DB.Model(&cloud.FileTrans{})
	var fileTranss []cloud.FileTrans
	var fileShares []cloud.FileShare
	// 查询出所有的id对应的文件uuid
	// 根据uuid生成/hls/uuid/index.m3u8  /hls/uuid/index.jpg   对应的key？暂不加密和ts 都需要添加
	// 后面加上 domain=%2A&expires=100&sign=zsK&timestamp=1667390405

	err = db.Where("id in ?", share.Ids).Find(&fileTranss).Error
	if err != nil {
		return
	}
	// 当前时间戳
	t := time.Now().Unix()
	// fmt.Println("t:", t)
	// 对选中的每一个视频制作m3u8和动态链接

	for _, v := range fileTranss {
		fileShare := cloud.FileShare{}
		uuid := v.TransUuid
		fileShare.FileName = v.TransOutName
		// path_m3u8 := "/hls/" + uuid + "/index.m3u8" //这个地址保持不变
		path_1080 := "/hls/" + uuid + "/1080p" //动态m3u8
		path_image := "/hls/" + uuid + "/index.jpg"
		// fileShare.M3Url = path_m3u8 + fileTransService.GenerateUrlTail(path_m3u8, t, share)

		fileTransService.GenerateNewM3u8(uuid, t, share) //生成新的m3u8文件
		fileShare.M3Url = path_1080 + fileTransService.GenerateUrlTail(path_1080, t, share)
		fileShare.PicUrl = path_image + fileTransService.GenerateUrlTail(path_image, t, share)
		fileShares = append(fileShares, fileShare)

	}

	return fileShares, err
}

// 生成URL后面的加密字符串
func (fileTransService *FileTransService) GenerateUrlTail(
	path string, t int64, share request.IdsShareReq) (tail string) {

	// /hls/bfd82b6a-b6f3-4433-8fe6-d8a928dfea45/index.m3u8 ?1667390167 ?100
	// 使用path?timesmap?expires  进行aes+base64加密作为sign传输
	// 接收方解密后，获得path 和请求的path 对比，不一致则退出；获取expires（秒）是否为-1，为-1则通过(有漏洞）。
	// 不为-1则计算 现在服务器时间-timesmap >expires 退出，否则返回视频数据
	// path:="/hls/"+uuid+"/index.m3u8"
	str_t := strconv.FormatInt(t, 10) //时间戳转字符串
	aes_str := path + "?" + str_t + "?" + share.Expires
	aes_sign, _ := tools.AesEcpt.AesBase64Encrypt(aes_str) //加密的sign
	tem_tail := "domain=" + share.Domain + "&expires=" + share.Expires +
		"&sign=" + url.QueryEscape(aes_sign) + "&timestamp=" + str_t
	tail = "?" + tem_tail
	// tail = "?" + tem_tail
	return tail
}

// 生成新的m3u8文件 动态生成
func (fileTransService *FileTransService) GenerateNewM3u8(uuid string,
	t int64, share request.IdsShareReq) {
	// base_path 系统存放m3u8的目录
	base_path := "D:\\myWork\\korea\\cloud\\translate\\transcoding-server\\temp\\testvideo\\out\\a\\"
	path_index_m3u8 := filepath.Join(base_path, uuid, "index.m3u8")
	path_1080_m3u8 := filepath.Join(base_path, uuid, "1080p")

	// 读取index.m3u8 同时生成动态加密tail
	fileHanle, err := os.OpenFile(path_index_m3u8, os.O_RDONLY, 0666)
	if err != nil {
		return
	}

	defer fileHanle.Close()
	// defer fileHanle1.Close()
	reader := bufio.NewReader(fileHanle)
	new_string := []string{}
	// 按行处理txt
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		m3 := string(line)
		if strings.HasPrefix(m3, "#") {
			// 没有加密的m3u8 不处理带key的
			new_string = append(new_string, m3)
			continue
		}
		// 计算tial
		tail := fileTransService.GenerateUrlTail("/hls/"+uuid+"/"+m3, t, share)
		m3 = m3 + tail
		new_string = append(new_string, m3)

	}

	tools.GenerateNewM3u8(new_string, path_1080_m3u8)

}
