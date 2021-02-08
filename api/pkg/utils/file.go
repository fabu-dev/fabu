package utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

//判断文件或文件夹是否存在 error为nil 则存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

// 获取文件的md5值
func FileHash(data []byte) string {
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}
