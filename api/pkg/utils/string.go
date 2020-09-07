package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

// 全局的随机种子
var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// 字符串拼接
func StringSplice(strList ...string) string {
	if len(strList) == 0 {
		return ""
	}

	var builder strings.Builder
	for _, str := range strList {
		builder.WriteString(str)
	}

	return builder.String()
}

//生成32位md5字串
func GetMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成随机字符串
func GetRandom(size int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, 0, size)
	for i := 0; i < size; i++ {
		result = append(result, str[r.Intn(len(str))])
	}

	return string(result)
}

//首字母转大写
func FirstLitterToUpper(str string) string {
	if len(str) == 0 {
		return ""
	}

	list := []rune(str)
	if list[0] >= 97 && list[0] <= 122 {
		list[0] -= 32
	}

	return string(list)
}
