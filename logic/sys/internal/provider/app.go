package provider

import (
	"github.com/fengjx/lucky/logic/sys/internal/service"
)

var AppProvider = &appProvider{}

type appProvider struct {
}

func (p appProvider) Init() error {
	return service.AppSvc.Init()
}
