package admin

import (
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/logic/cms/internal/admin/news"
)

func Init(hs *luchen.HTTPServer) {
	news.Init(hs)
}
