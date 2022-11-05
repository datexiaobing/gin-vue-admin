package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// 填充数据
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src) % blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

// 去掉填充数据
func unpadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}
func EncodePath(src []byte) (string, error) {
	return EncryptAES(src,[]byte("1233211233211233"))
}
func DecodePath(src []byte) (string, error)  {
	return DecryptAES(src,[]byte("1233211233211233"))
}

// 加密
func EncryptAES(src []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return  base64.StdEncoding.EncodeToString(src), nil
}

// 解密
func DecryptAES(src []byte, key []byte) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(string(src))
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(decoded, decoded)
	decoded = unpadding(decoded)
	return string(decoded), nil
}

