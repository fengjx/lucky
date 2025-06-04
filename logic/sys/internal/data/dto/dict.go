package dto

import (
	"github.com/fengjx/lucky/logic/sys/internal/dao/schema"
	"github.com/fengjx/lucky/logic/sys/syspub"
)

func BuildDictDTO(e *schema.SysDict) *syspub.DictDTO {
	if e == nil {
		return nil
	}
	return &syspub.DictDTO{
		Group: e.Group,
		Label: e.Label,
		Value: e.Value,
	}
}
