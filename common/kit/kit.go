package kit

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hash 生成 MD5 字符串
func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
