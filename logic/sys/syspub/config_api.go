package syspub

import (
	"github.com/fengjx/luchen/log"

	"github.com/fengjx/lucky/logic/sys/internal/data/consts"
)

const (
	// ScopeApp 应用级配置
	ScopeApp = consts.ScopeApp

	// ScopeSys 系统级配置
	ScopeSys = consts.ScopeSys
)

var ConfigAPI configAPI

type configAPI interface {
	// GetConfigString 返回key对应的配置
	GetConfigString(scope string, key string) string

	// GetConfig 返回key对应的配置，并序列化成对象
	GetConfig(scope string, key string, data any) error
}

func SetConfigAPI(impl configAPI) {
	if ConfigAPI != nil {
		log.Warn("set ConfigAPI impl duplicate")
		return
	}
	ConfigAPI = impl
}
