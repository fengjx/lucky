package provider

import (
	"context"

	"github.com/fengjx/lucky/logic/sys/internal/service"
)

var AppProvider = &appProvider{}

type appProvider struct {
}

func (p appProvider) Init(ctx context.Context) error {
	return service.AppSvc.Init(ctx)
}
