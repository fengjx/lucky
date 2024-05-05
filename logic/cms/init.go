package cms

import (
	"context"

	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/logic/cms/internal/endpoint"
)

func Init(ctx context.Context, httpServer *luchen.HTTPServer) {
	if httpServer != nil {
		endpoint.Init(ctx, httpServer)
	}
}
