package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
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
func CalcFileMD5(file *multipart.FileHeader) (string, error) {
	f, err := file.Open() //打开文件
	if nil != err {
		return "", err
	}
	defer f.Close()

	md5Handle := md5.New()         //创建 md5 句柄
	_, err = io.Copy(md5Handle, f) //将文件内容拷贝到 md5 句柄中
	if nil != err {
		return "", err
	}
	md := md5Handle.Sum(nil) //计算 MD5 值，返回 []byte

	md5str := fmt.Sprintf("%x", md) //将 []byte 转为 string
	return md5str, nil
}
