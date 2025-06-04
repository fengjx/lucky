package admin

import (
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/logic/sys/internal/admin/config"
	"github.com/fengjx/lucky/logic/sys/internal/admin/dict"
	"github.com/fengjx/lucky/logic/sys/internal/admin/user"
)

func Init(hs *luchen.HTTPServer) {
	user.Init(hs)
	config.Init(hs)
	dict.Init(hs)
}
