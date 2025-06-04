package cms

import (
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/logic/cms/internal/admin"
	"github.com/fengjx/lucky/logic/cms/internal/endpoint"
)

func Init(hs *luchen.HTTPServer) {
	admin.Init(hs)
	endpoint.Init(hs)
}
