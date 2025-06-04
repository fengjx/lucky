package endpoint

import (
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/logic/sys/internal/endpoint/app"
	"github.com/fengjx/lucky/logic/sys/internal/endpoint/login"
)

func Init(hs *luchen.HTTPServer) {
	login.RegisterLoginTTPHandler(hs)
	app.RegisterAppTTPHandler(hs)
}
