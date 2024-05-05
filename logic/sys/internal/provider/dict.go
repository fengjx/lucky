package provider

import (
	"github.com/fengjx/lucky/logic/sys/internal/service"
	"github.com/fengjx/lucky/logic/sys/syspub"
)

var DictProvider = &dictProvider{}

type dictProvider struct {
}

func (p dictProvider) GetGroupDict(group string) []*syspub.DictDTO {
	return service.DictSvc.GetGroupDict(group)[group]
}

func (p dictProvider) GetDictByLabel(group string, label string) *syspub.DictDTO {
	dictList := p.GetGroupDict(group)
	for _, dict := range dictList {
		if dict.Label == label {
			return dict
		}
	}
	return nil
}

func (p dictProvider) GetDictByValue(group string, value string) *syspub.DictDTO {
	dictList := p.GetGroupDict(group)
	for _, dict := range dictList {
		if dict.Value == value {
			return dict
		}
	}
	return nil
}
