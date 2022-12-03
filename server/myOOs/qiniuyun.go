package myoos

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

type QiNiu struct {
	Bucket    string
	AccessKey string
	SecretKey string
	QiNiuKey  string
}

func QiniuUpload(fileLocal string, uploadName string, qiniu QiNiu) (err error) {
	cfg := storage.Config{}

	// 空间对应的机房
	// cfg.Zone = &storage.ZoneXinjiapo
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// accessKey := "ROG6T97lMRotPTqwsUYin7g1SUovt9Hzyz-bJq4P"
	// secretKey := "u5ll7crp5o5Fsx154X44GXLdpfNJew3cei0wToDL"
	accessKey := qiniu.AccessKey
	secretKey := qiniu.SecretKey

	localFile := fileLocal
	bucket := qiniu.Bucket
	key := uploadName //上传后保存的文件名,包含路径

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "simsun ttc",
		},
	}
	err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(ret.Key, ret.Hash)
	return nil
}

// qiniu 到期时间16进制
func QiniuGetExpTime(exp_t time.Duration) string {
	t_now := time.Now()
	t_exp := t_now.Add(time.Second * exp_t).Unix() //默认30分钟过期
	t_16 := strconv.FormatInt(t_exp, 16)
	return t_16
}

func QiniuDeleteFile(qiniu QiNiu, qiniu_file string) error {

	// accessKey := "ROG6T97lMRotPTqwsUYin7g1SUovt9Hzyz-bJq4P"
	// secretKey := "u5ll7crp5o5Fsx154X44GXLdpfNJew3cei0wToDL"
	accessKey := qiniu.AccessKey
	secretKey := qiniu.SecretKey
	mac := auth.New(accessKey, secretKey)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	// cfg.Zone = &storage.ZoneXinjiapo
	bucketManager := storage.NewBucketManager(mac, &cfg)

	// key := "github.png"
	err := bucketManager.Delete(qiniu.Bucket, qiniu_file)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func QiniuReadIndexM3u8(path string, uuid string,
	qiniu_key string, exp_t string) (s []string, err error) {
	fileHanle, err := os.OpenFile(path, os.O_RDONLY, 0666)

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
			new_string = append(new_string, m3)
			continue
		}
		s := qiniu_key + "/video/hls/" + uuid + "/" + m3 + exp_t
		qiniu_sign := utils.MD5V([]byte(s))
		tail := "?sign=" + qiniu_sign + "&t=" + exp_t
		m3 = m3 + tail
		new_string = append(new_string, m3)

	}
	// fmt.Println(new_string)
	return new_string, nil
}
