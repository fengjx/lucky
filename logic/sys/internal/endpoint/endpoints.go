package endpoint

import (
	"context"

	"github.com/fengjx/luchen"
)

func Init(_ context.Context, httpServer *luchen.HTTPServer) {
	httpServer.Handler(
		&loginHandler{},
		&configHandler{},
		&configAdminHandler{},
		&userAdminHandler{},
		&sysDictAdminHandler{},
		&sysMenuAdminHandler{},
	)
}
