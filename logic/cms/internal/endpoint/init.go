package endpoint

import (
	"github.com/fengjx/luchen"
	"github.com/fengjx/lucky/logic/cms/internal/endpoint/news"
)

func Init(hs *luchen.HTTPServer) {
	news.RegisterNewsAdminTTPHandler(hs)

}
