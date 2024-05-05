package dto

import (
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/syspub"
)

func BuildConfigDTO(e *entity.SysConfig) *syspub.ConfigDTO {
	if e == nil {
		return nil
	}
	return &syspub.ConfigDTO{
		Scope: e.Scope,
		Key:   e.Key,
		Value: e.Value,
	}
}
