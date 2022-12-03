package cloud

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	cloudReq "github.com/flipped-aurora/gin-vue-admin/server/model/cloud/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	myoos "github.com/flipped-aurora/gin-vue-admin/server/myOOs"
	"github.com/flipped-aurora/gin-vue-admin/server/tools"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
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
	// 转换完成，删除原视频
	err = os.Remove(fileTrans.DownloadPath)
	return err
}

func (fileTransService *FileTransService) CreateFileTrans(fileTrans request.TransVideoReq) (err error) {

	var file cloud.FileTrans
	// linux 环境需要改这里
	baseOutPath := "/home/cloud/m3u8/nokey/"
	sysType := runtime.GOOS
	if sysType == "windows" {
		fmt.Println("Windows system")
		baseOutPath = "D:\\myWork\\korea\\cloud\\translate\\transcoding-server\\temp\\testvideo\\out\\a\\"
	}
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

// 点击分享，获取链接
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
	//先判断上次获取链接的时间距离现在是否超过半小时？
	for _, v := range fileTranss {
		fileShare := cloud.FileShare{}
		uuid := v.TransUuid
		fileShare.FileName = v.TransOutName
		// path_m3u8 := "/hls/" + uuid + "/index.m3u8" //这个地址保持不变
		path_1080 := "/hls/" + uuid + "/1080p.m3u8" //动态m3u8
		path_image := "/hls/" + uuid + "/index.jpg"
		// fileShare.M3Url = path_m3u8 + fileTransService.GenerateUrlTail(path_m3u8, t, share)
		fileTransService.GenerateNewM3u8(uuid, t, share) //生成新的m3u8文件
		fileShare.M3Url = path_1080 + fileTransService.GenerateUrlTail(path_1080, t, share)
		fileShare.PicUrl = path_image + fileTransService.GenerateUrlTail(path_image, t, share)

		// 获取系统配置
		var sysConfig cloud.CloudConfig
		err := global.GVA_DB.First(&sysConfig).Error
		if err != nil {
			fmt.Println("no  config information,ali url is null")
			continue
		}

		// 查看七牛状态是否为3，只有已上传到七牛的再获取七牛防盗链
		// fmt.Println("qiniu status:", v.TransOssQiniuStatus)
		if v.TransOssQiniuStatus == 3 {
			qiniu_base_url := "http://api.opkakao.com"
			qiniu_path := "video/hls/" + uuid + "/indexq.m3u8"      //七牛上的动态文件
			qiniu_key := "b68bbea9219bc8aa778a4dab2d8b2f89eb9a45a8" //备用key:ffb8feba2288f2aa16993588c58adfabbdb9adaa
			// bucket := "8886media"
			var qiniu myoos.QiNiu
			qiniu.AccessKey = "ROG6T97lMRotPTqwsUYin7g1SUovt9Hzyz-bJq4P"
			qiniu.SecretKey = "u5ll7crp5o5Fsx154X44GXLdpfNJew3cei0wToDL"
			//-------这两个会变----
			qiniu.Bucket = "8886media"
			qiniu.QiNiuKey = "b68bbea9219bc8aa778a4dab2d8b2f89eb9a45a8" //和bucket相关，设置bucket时新建
			// ex_t, _ := strconv.ParseInt(share.Expires, 10, 64)
			// var ex_t_qiniu int64
			// ex_t_qiniu = 1800
			// if ex_t > 0 {
			// 	ex_t_qiniu = ex_t
			// }
			exp_t := myoos.QiniuGetExpTime(1800) //默认30分钟有效
			s := qiniu_key + "/" + qiniu_path + exp_t
			// fmt.Println("s", s)
			qiniu_sign := utils.MD5V([]byte(s))
			qiniu_url := qiniu_base_url + "/" + qiniu_path + "?sign=" + qiniu_sign + "&t=" + exp_t
			fileShare.QiniuUrl = qiniu_url
			// 删除已有的uuid/index.m3u8 重新上传新的m3u8 带加密链接后缀的
			// to_path := filepath.Join(uuid, "index.m3u8")
			// to_s := filepath.ToSlash(to_path)
			err := myoos.QiniuDeleteFile(qiniu, qiniu_path) //删除m3u8
			if err != nil {
				fmt.Println("删除原来的m3u8失败", err)
			}
			// 新生成m3u8
			// base_path 系统存放m3u8的目录
			base_path := "/home/cloud/m3u8/nokey/"
			sysType := runtime.GOOS
			if sysType == "windows" {
				fmt.Println("Windows system")
				base_path = "D:\\myWork\\korea\\cloud\\translate\\transcoding-server\\temp\\testvideo\\out\\a\\"
			}
			new_qiniu_path := filepath.Join(base_path, uuid, "qiniu.m3u8")
			old_m3_path := filepath.Join(base_path, uuid, "index.m3u8")
			new_list, _ := myoos.QiniuReadIndexM3u8(old_m3_path, uuid, qiniu_key, exp_t)
			tools.GenerateNewM3u8(new_list, new_qiniu_path)
			// 上传新的qiniu.m3u8文件
			// new_qiniu_os_path := uuid + "/index.m3u8"
			err = myoos.QiniuUpload(new_qiniu_path, qiniu_path, qiniu)
			if err != nil {
				fmt.Println(err)
			}

		}

		// 查看阿里状态是否为3，只有已上传到阿里的再获取防盗链
		if v.TransOssAliStatus == 3 {
			// uuid := "f2473f93-f40f-4240-9437-a053e21c30e8"
			ali_save_base_path := "video/hls/" + uuid + "/"
			// 系统配置
			var aliParams myoos.AliParams
			// aliParams.Endpoint = "https://oss-cn-chengdu.aliyuncs.com"
			// aliParams.AccessKeyId = "LTAI5tFicacwP29FkJ2YNXqh"
			// aliParams.AccessKeySecret = "OGGweiUxSrD3Ub2dXpkMeyfB3tSTft"
			// aliParams.BucketName = "myceshialiyun"

			aliParams.Endpoint = sysConfig.AliEndPoint
			aliParams.AccessKeyId = sysConfig.AliAccessKeyId
			aliParams.AccessKeySecret = sysConfig.AliAccessKeySecret
			aliParams.BucketName = sysConfig.AliBucketName

			aliParams.Expt = 1800
			// 生成新的签名后的m3u8
			// base_path 系统存放m3u8的目录
			base_path := "/home/cloud/m3u8/nokey/"
			sysType := runtime.GOOS
			if sysType == "windows" {
				fmt.Println("Windows system")
				base_path = "D:\\myWork\\korea\\cloud\\translate\\transcoding-server\\temp\\testvideo\\out\\a\\"
			}
			new_ali_path := filepath.Join(base_path, uuid, "ali.m3u8")
			old_m3_path := filepath.Join(base_path, uuid, "index.m3u8")
			new_list, _ := myoos.AliReadIndexM3u8(old_m3_path, ali_save_base_path, aliParams)
			tools.GenerateNewM3u8(new_list, new_ali_path)
			// 上传新的qiniu.m3u8文件
			myoos.AliUpload(aliParams, new_ali_path, ali_save_base_path+"index.m3u8")
			// 获取新的m3u8加密链接
			new_ali_m3u8_singUrl := myoos.SignUrl(aliParams, ali_save_base_path+"index.m3u8")
			// fmt.Println(new_ali_m3u8_singUrl)
			fileShare.AliUrl = new_ali_m3u8_singUrl

		}

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
	base_path := "/home/cloud/m3u8/nokey/"
	sysType := runtime.GOOS
	if sysType == "windows" {
		fmt.Println("Windows system")
		base_path = "D:\\myWork\\korea\\cloud\\translate\\transcoding-server\\temp\\testvideo\\out\\a\\"
	}
	path_index_m3u8 := filepath.Join(base_path, uuid, "index.m3u8")
	path_1080_m3u8 := filepath.Join(base_path, uuid, "1080p.m3u8")

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

// 上传到七牛云
func (fileTransService *FileTransService) UploadQiniu(file cloud.FileTrans) (err error) {

	go fileTransService.UpToQiniu(file)
	// 此处修改状态为正在上传...1等待2正在3已上传
	err = global.GVA_DB.Model(cloud.FileTrans{}).Where("id =?", file.ID).Update("trans_oss_qiniu_status", 2).Error
	return err
}

func (fileTransService *FileTransService) UpToQiniu(file cloud.FileTrans) {
	// 读取对应uuid下的所有文件，依次上传到七牛
	// base_path 系统存放m3u8的目录
	base_path := "/home/cloud/m3u8/nokey/"
	sysType := runtime.GOOS
	if sysType == "windows" {
		fmt.Println("Windows system")
		base_path = "D:\\myWork\\korea\\cloud\\translate\\transcoding-server\\temp\\testvideo\\out\\a\\"
	}
	uuid := file.TransUuid
	video_path := filepath.Join(base_path, uuid)
	files, _ := ioutil.ReadDir(video_path)
	// qiniuBucktet := "8886media"
	var qiniu myoos.QiNiu
	qiniu.AccessKey = "ROG6T97lMRotPTqwsUYin7g1SUovt9Hzyz-bJq4P"
	qiniu.SecretKey = "u5ll7crp5o5Fsx154X44GXLdpfNJew3cei0wToDL"
	//-------这两个会变----
	qiniu.Bucket = "8886media"
	qiniu.QiNiuKey = "b68bbea9219bc8aa778a4dab2d8b2f89eb9a45a8"
	for _, f := range files {
		// fmt.Println(f.Name())
		from_path := filepath.Join(video_path, f.Name())
		form_s := filepath.ToSlash(from_path)
		to_path := filepath.Join("video/hls", uuid, f.Name())
		to_s := filepath.ToSlash(to_path)
		// fmt.Println(form_s, to_s, qiniuBucktet)
		myoos.QiniuUpload(form_s, to_s, qiniu)

	}
	// 更新七牛oss为已上传
	global.GVA_DB.Model(cloud.FileTrans{}).Where("id =?", file.ID).Update("trans_oss_qiniu_status", 3)
}

// 上传到ali oss
func (fileTransService *FileTransService) UploadAli(file cloud.FileTrans) (err error) {

	go fileTransService.UpToAli(file)
	// 此处修改状态为正在上传...1等待2正在3已上传
	err = global.GVA_DB.Model(cloud.FileTrans{}).Where("id =?", file.ID).Update("trans_oss_ali_status", 2).Error
	return err
}

func (fileTransService *FileTransService) UpToAli(file cloud.FileTrans) {

	// 系统配置 获取===
	var aliParams myoos.AliParams
	// aliParams.Endpoint = "https://oss-cn-chengdu.aliyuncs.com"
	// aliParams.AccessKeyId = "LTAI5tFicacwP29FkJ2YNXqh"
	// aliParams.AccessKeySecret = "OGGweiUxSrD3Ub2dXpkMeyfB3tSTft"
	// aliParams.BucketName = "myceshialiyun"

	var sysConfig cloud.CloudConfig

	err := global.GVA_DB.First(&sysConfig).Error
	if err != nil {
		fmt.Println("no  config information")
		return
	}
	aliParams.Endpoint = sysConfig.AliEndPoint
	aliParams.AccessKeyId = sysConfig.AliAccessKeyId
	aliParams.AccessKeySecret = sysConfig.AliAccessKeySecret
	aliParams.BucketName = sysConfig.AliBucketName

	aliParams.Expt = 1800
	// base_path 系统存放m3u8的目录
	base_path := "/home/cloud/m3u8/nokey/"
	sysType := runtime.GOOS
	if sysType == "windows" {
		fmt.Println("Windows system")
		base_path = "D:\\myWork\\korea\\cloud\\translate\\transcoding-server\\temp\\testvideo\\out\\a\\"
	}
	uuid := file.TransUuid
	video_path := filepath.Join(base_path, uuid) //本地文件位置
	ali_save_base_path := "video/hls/" + uuid + "/"
	files, _ := ioutil.ReadDir(video_path)

	for _, f := range files {
		from_path := filepath.Join(video_path, f.Name())
		form_s := filepath.ToSlash(from_path)
		to_path := ali_save_base_path + f.Name()
		// fmt.Println(form_s, to_s, qiniuBucktet)
		myoos.AliUpload(aliParams, form_s, to_path)

	}
	// 更新阿里oss为已上传
	global.GVA_DB.Model(cloud.FileTrans{}).Where("id =?", file.ID).Update("trans_oss_ali_status", 3)
}
