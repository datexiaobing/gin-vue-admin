package utils

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type PostBody struct {
	Uuid  string `json:"uuid"`
	Title string `json:"title"`
	OssId int    `json:"ossId"`
	// Video       string `json:"video"` //视频链接
	// VideoCover  string `json:"videoCover"`
	// VideoCover2 string `json:"video_cover_2 "`
	ClassifyId int    `json:"classifyId"`
	Duration   string `json:"duration"`
	SpecialId  int    `json:"specialId"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Post("http://xxxx","application/json;charset=utf-8",[]byte("{'aaa':'bbb'}"))
func GoPost(url string, contentType string, body []byte) (string, error) {
	res, err := http.Post(url, contentType, strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
