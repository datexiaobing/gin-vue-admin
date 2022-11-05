package tools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func Md5toString(f string) string {
	w := md5.New()
	io.WriteString(w, f)
	return fmt.Sprintf("%x", w.Sum(nil))
}

//返回一个32位md5加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}


func Get16MD5Encode(data string) string{
	return GetMD5Encode(data)[8:24]
}