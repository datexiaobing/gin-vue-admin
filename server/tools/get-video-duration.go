package tools

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

//constant.DurationFormat="00:00:00"

// BoxHeader 信息头
type BoxHeader struct {
	Size       uint32
	FourccType [4]byte
	Size64     uint64
}

// filePath 视频地址
func GetMP4Duration(filePath string) (duration string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return "00:00:00", errors.New("文件不存在")
		}
		return "00:00:00", err
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			zap.L().Error("file close fail err:", zap.Error(err))
		}
	}(file)

	var (
		info      = make([]byte, 0x10)
		boxHeader BoxHeader
		offset    int64 = 0
	)
	// 获取结构偏移
	for {
		_, err = file.ReadAt(info, offset)
		if err != nil {
			return
		}
		boxHeader = getHeaderBoxInfo(info)
		fourccType := getFourccType(boxHeader)
		if fourccType == "moov" {
			break
		}
		// 有一部分mp4 mdat尺寸过大需要特殊处理
		if fourccType == "mdat" {
			if boxHeader.Size == 1 {
				offset += int64(boxHeader.Size64)
				continue
			}
		}
		offset += int64(boxHeader.Size)
	}
	// 获取move结构开头一部分
	moveStartBytes := make([]byte, 0x100)
	_, err = file.ReadAt(moveStartBytes, offset)
	if err != nil {
		return
	}
	// 定义timeScale与Duration偏移
	timeScaleOffset := 0x1C
	durationOffset := 0x20
	timeScale := binary.BigEndian.Uint32(moveStartBytes[timeScaleOffset : timeScaleOffset+4])
	Duration := binary.BigEndian.Uint32(moveStartBytes[durationOffset : durationOffset+4])
	return ResolveTime(Duration / timeScale), nil
}

// getHeaderBoxInfo 获取头信息
func getHeaderBoxInfo(data []byte) (boxHeader BoxHeader) {
	buf := bytes.NewBuffer(data)
	_ = binary.Read(buf, binary.BigEndian, &boxHeader)
	return
}

// getFourccType 获取信息头类型
func getFourccType(boxHeader BoxHeader) (fourccType string) {
	fourccType = string(boxHeader.FourccType[:])
	return
}

// ResolveTime 将秒转成时分秒格式
func ResolveTime(seconds uint32) string {
	var (
		h, m, s string
	)
	var day = seconds / (24 * 3600)
	hour := (seconds - day*3600*24) / 3600
	minute := (seconds - day*24*3600 - hour*3600) / 60
	second := seconds - day*24*3600 - hour*3600 - minute*60
	h = strconv.Itoa(int(hour))
	if hour < 10 {
		h = "0" + strconv.Itoa(int(hour))
	}
	m = strconv.Itoa(int(minute))
	if minute < 10 {
		m = "0" + strconv.Itoa(int(minute))
	}
	s = strconv.Itoa(int(second))
	if second < 10 {
		s = "0" + strconv.Itoa(int(second))
	}
	return fmt.Sprintf("%s:%s:%s", h, m, s)
}
