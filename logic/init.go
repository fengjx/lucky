package logic

import (
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/common/lifecycle"
	"github.com/fengjx/lucky/logic/cms"
	"github.com/fengjx/lucky/logic/sys"
)

func Init(hs *luchen.HTTPServer) {
	sys.Init(hs)
	cms.Init(hs)
	lifecycle.DoHooks()
}
