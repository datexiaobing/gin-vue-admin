package tools

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 没加密key的m3u8
func ReadIndexM3u8(path string, tail string) (s []string, err error) {
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
		m3 = m3 + tail
		new_string = append(new_string, m3)

	}
	// fmt.Println(new_string)
	return new_string, nil
}

func GenerateNewM3u8(s []string, path string) (err error) {
	fileHanle, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fileHanle.Close()
	w := bufio.NewWriter(fileHanle)
	var new_s string
	for _, v := range s {
		new_s += v + "\n"
	}
	w.WriteString(new_s)
	w.Flush()

	return nil
}
