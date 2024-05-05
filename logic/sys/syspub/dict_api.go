package syspub

import "github.com/fengjx/luchen/log"

var DictAPI dictAPI

type dictAPI interface {

	// GetGroupDict 根据分组获取字典列表
	GetGroupDict(group string) []*DictDTO

	// GetDictByLabel 根据分组和标签获取字典
	GetDictByLabel(group string, label string) *DictDTO

	// GetDictByValue 根据分组和值获取字典
	GetDictByValue(group string, label string) *DictDTO
}

func SetDictAPI(impl dictAPI) {
	if DictAPI != nil {
		log.Warn("set DictAPI impl duplicate")
		return
	}
	DictAPI = impl
}
