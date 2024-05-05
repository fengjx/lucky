package sys

import (
	"context"
	"time"

	"github.com/fengjx/go-halo/hook"
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/connom/lifecycle"

	"github.com/fengjx/lucky/logic/sys/internal/endpoint"
	"github.com/fengjx/lucky/logic/sys/internal/provider"
	"github.com/fengjx/lucky/logic/sys/internal/service"
	"github.com/fengjx/lucky/logic/sys/syspub"
)

// Init 初始化
func Init(ctx context.Context, httpServer *luchen.HTTPServer) {
	lifecycle.AddHook(lifecycle.InterfaceAware, func() {
		syspub.SetDictAPI(provider.DictProvider)
		syspub.SetConfigAPI(provider.ConfigProvider)
		syspub.SetAppAPI(provider.AppProvider)
	})
	lifecycle.AddHook(lifecycle.InitData, func() {
		service.ConfigSvc.Refresh(ctx)
	}, hook.WithInterval(time.Minute))

	lifecycle.AddHook(lifecycle.InitData, func() {
		service.DictSvc.Refresh(ctx)
	}, hook.WithInterval(time.Minute))

	if httpServer != nil {
		endpoint.Init(ctx, httpServer)
	}
}

// InitWithTools 执行工具脚本时使用的初始化逻辑
func InitWithTools(ctx context.Context) {
	syspub.SetDictAPI(provider.DictProvider)
	syspub.SetConfigAPI(provider.ConfigProvider)
	syspub.SetAppAPI(provider.AppProvider)
}
