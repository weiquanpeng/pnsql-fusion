package utils

import (
	"crypto/md5"
	"encoding/hex"
)

var (
	PWD_SALT = "pengweiquan1021555x3"
)

// 参数：需要加密的字符串
func getMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// md5加密
func Md5(str string) string {
	return getMd5(getMd5(PWD_SALT + str + PWD_SALT))
}

// md5 加盐
//func Md5Slat(str string, slat string) string {
//	return getMd5(getMd5(PWD_SALT + str + slat))
//}

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
