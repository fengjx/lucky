package dto

import (
	"github.com/fengjx/lucky/logic/sys/internal/dao/schema"
	"github.com/fengjx/lucky/logic/sys/syspub"
)

func BuildConfigDTO(e *schema.SysConfig) *syspub.ConfigDTO {
	if e == nil {
		return nil
	}
	return &syspub.ConfigDTO{
		Scope: e.Scope,
		Key:   e.Key,
		Value: e.Value,
	}
}
