package config

import (
	"time"

	"github.com/fengjx/go-halo/json"
)

func init() {
	// time 类型转换为时间戳
	json.RegisterTimeAsInt64Codec(time.Millisecond)
	// 允许字符串和数值类型互转
	json.RegisterFuzzyDecoders()
}
