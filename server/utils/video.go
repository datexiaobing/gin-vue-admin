package utils

import (
	"fmt"
	"strings"

	aria2go "github.com/flipped-aurora/gin-vue-admin/server/aria2"
)

func SplitUrls(s string) []string {
	return strings.Split(s, "\n")
}

// var client *aria2go.Aria2Client

// Download 下载文件
func DownloadByUrl(url string) (gid string) {
	client := aria2go.NewAria2Client("8888")
	gid, err := client.Download(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return gid
}

// active video
func QueryActiveVideos() (task []*aria2go.TaskStatusData, err error) {
	client := aria2go.NewAria2Client("8888")
	activeTask, err := client.QueryDownloadingTask()
	return activeTask, err
}

// waiting video
func QueryWaitingVideos(offset int, limit int) (task []*aria2go.TaskStatusData, err error) {
	client := aria2go.NewAria2Client("8888")
	task, err = client.QueryWaitingTask(offset, limit)
	return task, err
}

// stoped video
func QueryStopVideos(offset int, limit int) (task []*aria2go.TaskStatusData, err error) {
	client := aria2go.NewAria2Client("8888")
	task, err = client.QueryStoppedTask(offset, limit)
	return task, err
}

// pause video
func PauseVideo(gid string) (err error) {
	client := aria2go.NewAria2Client("8888")
	err = client.Pause(gid)
	return err
}

// unpause video
func UnpauseVideo(gid string) (err error) {
	client := aria2go.NewAria2Client("8888")
	err = client.Unpause(gid)
	return err
}

// remove task
func RemoveTask(gid string) (err error) {
	client := aria2go.NewAria2Client("8888")
	err = client.RemoveTask(gid)
	return err
}

func DownloadTorrent(path string) (gid string, err error) {
	client := aria2go.NewAria2Client("8888")
	gid, err = client.DownloadWithLocalTorrent(path)
	return gid, err
}
