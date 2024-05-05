package logic

import (
	"context"

	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/connom/lifecycle"
	"github.com/fengjx/lucky/logic/cms"
	"github.com/fengjx/lucky/logic/sys"
)

func Init(ctx context.Context, httpServer *luchen.HTTPServer) {
	sys.Init(ctx, httpServer)
	cms.Init(ctx, httpServer)
	lifecycle.DoHooks()
}
