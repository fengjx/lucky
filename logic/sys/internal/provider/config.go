package provider

import (
	"github.com/fengjx/go-halo/json"

	"github.com/fengjx/lucky/logic/sys/internal/service"
)

var ConfigProvider = &configProvider{}

type configProvider struct {
}

func (p configProvider) GetConfigString(scope string, key string) string {
	configList := service.ConfigSvc.ScopeConfig(scope)[scope]
	for _, config := range configList {
		if config.Key == key {
			return config.Value
		}
	}
	return ""
}

func (p configProvider) GetConfig(scope string, key string, data any) error {
	value := p.GetConfigString(scope, key)
	if value == "" {
		return nil
	}
	return json.FromJson(value, data)
}
