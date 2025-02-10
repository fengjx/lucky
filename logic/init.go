package logic

import (
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/logic/sys"
	"github.com/fengjx/lucky/pkg/lifecycle"
)

func Init(httpServer *luchen.HTTPServer) {
	sys.Init(httpServer)
	lifecycle.DoHooks()
}
