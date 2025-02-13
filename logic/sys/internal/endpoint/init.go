package endpoint

import (
	"github.com/fengjx/luchen"
	"github.com/fengjx/lucky/logic/sys/internal/endpoint/app"
	"github.com/fengjx/lucky/logic/sys/internal/endpoint/config"
	"github.com/fengjx/lucky/logic/sys/internal/endpoint/dict"
	"github.com/fengjx/lucky/logic/sys/internal/endpoint/login"
	"github.com/fengjx/lucky/logic/sys/internal/endpoint/menu"
	"github.com/fengjx/lucky/logic/sys/internal/endpoint/user"
)

func Init(hs *luchen.HTTPServer) {
	login.RegisterLoginTTPHandler(hs)
	app.RegisterAppTTPHandler(hs)
	config.RegisterConfigAdminTTPHandler(hs)
	user.RegisterUserAdminEndpoint(hs)
	dict.RegisterDictAdminTTPHandler(hs)
	menu.RegisterMenuAdminTTPHandler(hs)
}
