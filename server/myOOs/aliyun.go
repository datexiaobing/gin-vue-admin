package myoos

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliParams struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	// ObjectName      string
	Expt int64
}

func HandleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

// 上传文件、存在时则覆盖
func AliUpload(aliParams AliParams, localFileName string, objectName string) {
	// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。
	// objectName := "video/hls/kjkjl-uuik/playlist.m3u8"
	// yourLocalFileName填写本地文件的完整路径。
	// localFileName := base_path + "/playlist.m3u8"
	// 创建OSSClient实例。
	client, err := oss.New(aliParams.Endpoint, aliParams.AccessKeyId, aliParams.AccessKeySecret)
	if err != nil {
		HandleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(aliParams.BucketName)
	if err != nil {
		HandleError(err)
	}

	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		HandleError(err)
	}
}

func SignUrl(aliParams AliParams, ObjectName string) (signalUrl string) {
	// endpoint := "https://oss-cn-chengdu.aliyuncs.com"
	// accessKeyId := "LTAI5tFicacwP29FkJ2YNXqh"
	// accessKeySecret := "OGGweiUxSrD3Ub2dXpkMeyfB3tSTft"
	// yourBucketName填写存储空间名称。
	// bucketName := "myceshialiyun"
	// 填写文件完整路径，例如exampledir/exampleobject.txt。文件完整路径中不能包含Bucket名称。
	// objectName := "video/hls/f2473f93-f40f-4240-9437-a053e21c30e8/index1.ts"
	client, err := oss.New(aliParams.Endpoint, aliParams.AccessKeyId, aliParams.AccessKeySecret)
	if err != nil {
		HandleError(err)
	}

	bucket, err := client.Bucket(aliParams.BucketName)
	if err != nil {
		HandleError(err)
	}
	// 生成用于下载的签名URL，并指定签名URL的有效时间为60秒。
	signedURL, err := bucket.SignURL(ObjectName, oss.HTTPGet, aliParams.Expt)
	if err != nil {
		HandleError(err)
	}
	// fmt.Printf("Sign Url:%s\n", signedURL)
	return signedURL
}

func AliReadIndexM3u8(path string, ali_save_base_path string, aliParams AliParams) (s []string, err error) {
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
		ObjectName := ali_save_base_path + m3
		m3 = SignUrl(aliParams, ObjectName)
		new_string = append(new_string, m3)

	}
	// fmt.Println(new_string)
	return new_string, nil
}

//构建一个阿里云客户端, 用于发起请求。
//设置调用者（RAM用户或RAM角色）的AccessKey ID和AccessKey Secret。
// client1, _ := sts.NewClientWithAccessKey("cn-chengdu", accessKeyId, accessKeySecret)

// //构建请求对象。
// request := sts.CreateAssumeRoleRequest()
// request.Scheme = "https"

// //设置参数。
// rolArn := "acs:ram::1544644851108122:role/videoarm"
// request.RoleArn = rolArn
// request.RoleSessionName = "admin"

// //发起请求，并得到响应。
// response, err := client1.AssumeRole(request)
// if err != nil {
// 	fmt.Print(err.Error())
// }
// fmt.Printf("response is %#v\n", response)

// securityToken := response.Credentials.SecurityToken
// accessKeyId1 := response.Credentials.AccessKeyId
// accessKeySecret1 := response.Credentials.AccessKeySecret
// oss.SecurityToken(securityToken)
// 获取STS临时凭证后，您可以通过其中的安全令牌（SecurityToken）和临时访问密钥（AccessKeyId和AccessKeySecret）生成OSSClient。
