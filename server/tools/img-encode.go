/**
 * 图片加密解密工具类
 **/
package tools

import (
	"os"

	"github.com/thep0y/go-logger/log"
)

func ImgEncode(f string) ([]byte, error) {
	lisr, err := File2Bytes(f)
	if err != nil {
		log.Errorf("File2Bytes文件转二进制数组失败 error:%s \r\n", err)
		return nil, err
	}

	data := make([]byte, 0)
	for _, value := range lisr {
		y := int(value) - 1
		data = append(data, uint8(y))
	}

	return data, nil
}

func ImgDecode(spath string, epath string) error {
	lisr, err := File2Bytes(epath)
	if err != nil {
		log.Errorf("File2Bytes文件转二进制数组失败 error:%s \r\n", err)
		return err
	}

	data := make([]byte, 0)
	for _, value := range lisr {
		y := int(value) + 1
		data = append(data, uint8(y))
	}
	os.WriteFile(spath, data, 0666)
	return nil
}

func File2Bytes(filename string) ([]byte, error) {
	// File
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// FileInfo:
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// []byte
	data := make([]byte, stats.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
