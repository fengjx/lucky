package logic

import (
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/logic/cms"
	"github.com/fengjx/lucky/logic/sys"
	"github.com/fengjx/lucky/pkg/lifecycle"
)

func Init(hs *luchen.HTTPServer) {
	sys.Init(hs)
	cms.Init(hs)
	lifecycle.DoHooks()
}
