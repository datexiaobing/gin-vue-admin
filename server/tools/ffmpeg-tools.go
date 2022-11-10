package tools

import (
	//"go-admin/common/ffmpegManage/offlinedownload"
	"fmt"
	"image/png"
	"math"
	"os"
	"strconv"

	log "github.com/flipped-aurora/gin-vue-admin/server/go-admin-core/logger"
	"github.com/flipped-aurora/gin-vue-admin/server/go-admin-core/sdk/config"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cloud"
	"github.com/flipped-aurora/gin-vue-admin/server/tools/ffmpegBinding/transcoder"
	"github.com/flipped-aurora/gin-vue-admin/server/tools/ffmpegBinding/transcoder/ffmpeg"
)

// D:\myWork\ffmpeg\ffmpegmaster\bin
var FfmpegBinPath = "D:\\myWork\\ffmpeg\\ffmpegmaster\\bin\\ffmpeg.exe"
var FfprobeBinPath = "D:\\myWork\\ffmpeg\\ffmpegmaster\\bin\\ffprobe.exe"

// var FfmpegBinPath = "/usr/bin/ffmpeg"
// var FfprobeBinPath = "/usr/bin/ffprobe"

/*
*

	ffmpeg进度通知
*/
type WriteCounter interface {
	PrintProgress(n float64)
}

/*
*

	转码设置模型
*/
type TranscodingConfig struct {
	InputPath       string //需要转码的file路径
	OutputPath      string //转出存放的目录    不带文件名
	OutputType      string //转出文件格式      转出文件格式mp4 m3u8
	HlsTime         int    //单片时常单位秒
	HlsKey          bool   //是否需要切片加密
	ResolutionRatio int    // 0流畅 1普通 2高清 3超清
	StatrTime       string //开始时间
	Fps             int
}

/*
*

	水印设置模型
*/
type WatermarkConfig struct {
	Switch       bool   //开关
	Type         int    //水印类型  1固定文字 2图片logo
	ImgPath      string //图片路径
	Lutyuv       string //透明度
	Text         string //文字内容
	Fontsize     int    //字体大小
	Fontcolor    string //字体颜色
	Fontfile     string //字体路径
	TextType     int    //文字水印类型  0固定位子 1滚动
	XCoordinate  int
	YCoordinate  int
	BasicPoint   int //基点 0右下角 1右上角 2左下角 3左上角
	RollingStatr int //滚动开始时间 秒
	RollingSpace int //间隔5秒循环
	RollingSpeed int //滚动速度
}

/*
*
视频截图模型
*/
type VideoCaptureConfig struct {
	Switch    bool   //开关
	StatrTime string //开始时间
}

/*
*

	获取视频宽高
*/
func GetVideoWH(inputPath string) (int, int, error) {

	outputPathfile := "temp/" + Md5toString(inputPath) + ".png"
	fmt.Println("get videoWH:", outputPathfile)

	err := VideoCapture(inputPath, outputPathfile, VideoCaptureConfig{StatrTime: "1.0"})
	if err != nil {
		log.Error("GetVideoWH GetVideoWH file:", outputPathfile, " error:", err)
		return 0, 0, err
	}
	defer os.Remove(outputPathfile)

	file, err := os.Open(outputPathfile)
	defer file.Close()
	if err != nil {
		log.Error("GetVideoWH GetVideoWH file:", outputPathfile, " error:", err)
		return 0, 0, err
	}

	img, err := png.Decode(file)
	if err != nil {
		log.Error("GetVideoWH GetVideoWH file:", outputPathfile, " error:", err)
		return 0, 0, err
	}
	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	return width, height, err
}

/*
*

	视频截图
*/
func VideoCapture(inputPath string, outputPath string, videoCaptureConfig VideoCaptureConfig) error {
	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath:   FfmpegBinPath,
		FfprobeBinPath:  FfprobeBinPath,
		ProgressEnabled: true,
	}
	OutputFormat := "image2"
	SeekTime := videoCaptureConfig.StatrTime
	Duration := "0.001"
	Overwrite := true
	options := ffmpeg.Options{
		OutputFormat: &OutputFormat,
		SeekTime:     &SeekTime,
		Duration:     &Duration,
		Overwrite:    &Overwrite,
	}

	var transcoder transcoder.Information
	ffmpeg.
		New(ffmpegConf).
		Input(inputPath).
		Output(outputPath).
		WithOptions(options).
		Run(&transcoder)

	progress := transcoder.Progress
	done := transcoder.Error

	for range progress {
	}

	err := <-done
	if err != nil {
		log.Error("VideoCapture FFmpeg Run error:", err)
		return err
	}

	return nil

}

/*
*

	转码函数
*/
func Transcoding(tconfig TranscodingConfig,
	wconfig WatermarkConfig,
	cconfig VideoCaptureConfig,
	file cloud.FileTrans,
	write func(msg float64, file cloud.FileTrans)) error {

	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath:   FfmpegBinPath,
		FfprobeBinPath:  FfprobeBinPath,
		ProgressEnabled: true,
		Verbose:         config.Ffmpegconf.Verbose,
	}

	inputPath := tconfig.InputPath
	outputPath := tconfig.OutputPath
	outputType := tconfig.OutputType
	hlsTime := tconfig.HlsTime
	hlsKey := tconfig.HlsKey
	resolutionRatio := tconfig.ResolutionRatio
	statrTime := tconfig.StatrTime
	fps := tconfig.Fps

	_, err := os.Stat(outputPath)
	res := os.IsNotExist(err)
	if res == true {
		os.MkdirAll(outputPath, os.ModePerm)
	}

	if outputType == "m3u8" && hlsKey {
		if !PathExists(outputPath + "/index.key") {
			key := CreateRandomString(16)
			err := WriteFile(outputPath+"/index.key", key)
			if err != nil {
				log.Error("Transcoding WriteFile:", outputPath+"/index.key", " error:", err)
				return err
			}
			keyinfo := "" +
				"index.key \n" +
				outputPath + "/index.key"
			err = WriteFile(outputPath+"/index.keyinfo", keyinfo)
			if err != nil {
				log.Error("Transcoding WriteFile:", outputPath+"/index.keyinfo", " error:", err)
				return err
			}
		}
	}

	options := new(ffmpeg.Options)
	CodecV := "libx264"
	CodecA := "copy"
	Overwrite := true
	options.CodecV = &CodecV
	options.CodecA = &CodecA
	options.Overwrite = &Overwrite

	if wconfig.Switch {
		// 文字跑马灯
		if wconfig.Type == 1 {
			Fontcolor := wconfig.Fontcolor
			if Fontcolor == "" {
				Fontcolor = "white@1.0"
			}

			Fontsize := wconfig.Fontsize
			if Fontsize == 0 {
				Fontsize = 36
			}

			XCoordinate := strconv.Itoa(wconfig.XCoordinate)
			YCoordinate := strconv.Itoa(wconfig.YCoordinate)
			xy := "y=" + YCoordinate + ":x=" + XCoordinate

			if wconfig.TextType == 0 {
				if wconfig.BasicPoint == 0 {
					xy = "y=main_w-overlay_w-" + YCoordinate + ":x=main_h-overlay_h-" + XCoordinate
				} else if wconfig.BasicPoint == 1 {
					xy = "y=main_w-overlay_w-" + YCoordinate + ":x=" + XCoordinate
				} else if wconfig.BasicPoint == 2 {
					xy = "y=" + YCoordinate + ":x=main_h-overlay_h-" + XCoordinate
				} else if wconfig.BasicPoint == 3 {
					xy = "y=" + YCoordinate + ":x=" + XCoordinate
				}
			} else {
				RollingStatr := strconv.Itoa(wconfig.RollingStatr) //滚动开始时间 秒
				RollingSpace := strconv.Itoa(wconfig.RollingSpace) //间隔5秒循环
				RollingSpeed := strconv.Itoa(wconfig.RollingSpeed) //滚动速度
				if file.TransDrawtextPosition == 1 {
					// 上方
					xy = "y=" + YCoordinate + ":x=(w)+(" + RollingStatr + "*" + RollingSpeed + ") - mod((t*" + RollingSpeed + ")" + "\\" + ",tw+w+(" + RollingSpace + "*" + RollingSpeed + "))"
				} else {
					// 下方
					xy = "y=h-line_h" + ":x=(w)+(" + RollingStatr + "*" + RollingSpeed + ") - mod((t*" + RollingSpeed + ")" + "\\" + ",tw+w+(" + RollingSpace + "*" + RollingSpeed + "))"
				}
			}
			// rootpath := GetCurrentDirectory()
			VideoFilter := "drawtext=" +
				"text=" + "'" + wconfig.Text + "'" +
				":fontcolor=" + Fontcolor + "" +
				":fontsize=" + strconv.Itoa(Fontsize) + "" +
				":fontfile=./simsun.ttc" +
				":" + xy + "" +
				""
			options.VideoFilter = &VideoFilter

		} else if wconfig.Type == 2 {
			XCoordinate := strconv.Itoa(wconfig.XCoordinate)
			YCoordinate := strconv.Itoa(wconfig.YCoordinate)
			overlay := "overlay=0:0"
			if wconfig.BasicPoint == 0 {
				overlay = "overlay=main_w-overlay_w-" + YCoordinate + ":main_h-overlay_h-" + XCoordinate
			} else if wconfig.BasicPoint == 1 {
				overlay = "overlay=main_w-overlay_w-" + YCoordinate + ":" + XCoordinate
			} else if wconfig.BasicPoint == 2 {
				overlay = "overlay=" + YCoordinate + ":main_h-overlay_h-" + XCoordinate
			} else if wconfig.BasicPoint == 3 {
				overlay = "overlay=" + YCoordinate + ":" + XCoordinate
			}
			//scale= 50:50
			VideoFilter := "movie=" + wconfig.ImgPath + ",lutyuv=a=val*" + wconfig.Lutyuv + "[watermark];[in][watermark] " + overlay + "[out]"
			options.VideoFilter = &VideoFilter
		}

	}

	if outputType == "m3u8" {
		if hlsKey {
			EncryptionKey := outputPath + "/index.keyinfo"
			options.EncryptionKey = &EncryptionKey
		}

		if hlsTime == 0 {
			hlsTime = 3
		}
		options.HlsSegmentDuration = &hlsTime
		Strict := 2
		options.Strict = &Strict
		OutputFormat := "hls"
		options.OutputFormat = &OutputFormat
		HlsListSize := 0
		options.HlsListSize = &HlsListSize
		HlsPlaylistType := "vod"
		options.HlsPlaylistType = &HlsPlaylistType
		ForceKeyFrames := "expr:gte(t,n_forced*" + strconv.Itoa(hlsTime) + ")"
		options.ForceKeyFrames = &ForceKeyFrames

	}

	width, height, err := GetVideoWH(tconfig.InputPath)
	// fmt.Println(width)
	// width, height, err := 100, 100, nil
	if err != nil {
		return err
	}
	model := Model(int64(width), int64(height))

	ishengp := true
	if width < height {
		ishengp = false
	}

	if resolutionRatio > 0 {
		contrast := model.Verdict()
		videop := 360
		if resolutionRatio == 1 {
			videop = 360
		} else if resolutionRatio == 2 {
			videop = 720
		} else if resolutionRatio == 3 {
			videop = 1080
		} else if resolutionRatio == 4 {
			videop = 1440
		}

		if ishengp {
			height = videop
			width = int(math.Ceil(float64(height) * contrast))
			if width%2 != 0 {
				width = width + 1
			}
		} else {
			width = videop
			height = int(math.Ceil(float64(height) / contrast))
			if height%2 != 0 {
				height = height + 1
			}
		}

		Resolution := strconv.Itoa(width) + "x" + strconv.Itoa(height)
		options.Resolution = &Resolution
	}
	if statrTime != "" {
		options.SeekTime = &statrTime
	}

	if fps > 0 {
		options.FrameRate = &fps
	}

	outputPathfile := outputPath + "/index." + outputType
	// fmt.Println(outputPathfile)
	if PathExists(outputPathfile) {
		os.Remove(outputPathfile)
	}

	var transcoder transcoder.Information
	ffmpeg.
		New(ffmpegConf).
		Input(inputPath).
		Output(outputPathfile).
		WithOptions(options).
		Run(&transcoder)

	progress := transcoder.Progress
	// fmt.Println(progress)
	done := transcoder.Error

	go func() {
		for msg := range progress {
			Progress := msg.GetProgress()
			if write != nil {
				write(Progress, file)
			}
		}
	}()

	err = <-done
	if err == nil {
		if write != nil {
			write(99.99, file)
		}
	} else {
		log.Error("Transcoding FFmpeg Run error:", err)
		return err
	}
	if cconfig.Switch {
		// 截图
		err = VideoCapture(inputPath, outputPath+"/index.jpg", cconfig)
	}

	if err == nil {
		if write != nil {
			write(100.00, file)
		}
	}

	return err

}
