package endpoint

import (
	"github.com/fengjx/luchen"
)

func Init(httpServer *luchen.HTTPServer) {
	httpServer.Handler(
		&newsAdminHandler{},
	)
}
