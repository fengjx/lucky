package dto

import "github.com/fengjx/lucky/logic/sys/internal/data/entity"

type MenuDTO struct {
	*entity.SysMenu
	Children []*MenuDTO `json:"children"`
}

func BuildMenuDTO(e *entity.SysMenu) *MenuDTO {
	if e == nil {
		return nil
	}
	return &MenuDTO{
		SysMenu:  e,
		Children: nil,
	}
}
