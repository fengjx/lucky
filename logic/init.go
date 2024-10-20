package logic

import (
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/connom/lifecycle"
	"github.com/fengjx/lucky/logic/cms"
	"github.com/fengjx/lucky/logic/sys"
)

func Init(httpServer *luchen.HTTPServer) {
	sys.Init(httpServer)
	cms.Init(httpServer)
	lifecycle.DoHooks()
}
