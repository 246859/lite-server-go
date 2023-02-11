package ginutils

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// Sha1
// @Date 2023-02-11 20:08:10
// @Param str string
// @Return hash string
// @Method
// @Description: 将一个字符串以sha1进行编码
func Sha1(str string) (hash string) {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha256
// @Date 2023-02-11 20:09:35
// @Param str string
// @Return hash256 string
// @Method
// @Description: 将一个字符串以sha256加密
func Sha256(str string) (hash256 string) {
	hash := sha256.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
