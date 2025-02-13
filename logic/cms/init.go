package cms

import (
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/logic/cms/internal/endpoint"
)

func Init(httpServer *luchen.HTTPServer) {
	if httpServer != nil {
		endpoint.Init(httpServer)
	}
}
