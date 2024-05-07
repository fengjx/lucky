package config

import "github.com/fengjx/luchen/env"

// IsDemoEnv 是否是demo环境
func IsDemoEnv() bool {
	return env.Is("demo")
}
