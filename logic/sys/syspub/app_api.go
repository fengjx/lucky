package syspub

import (
	"context"

	"github.com/fengjx/luchen/log"
)

var AppAPI appAPI

type appAPI interface {
	Init(ctx context.Context) error
}

func SetAppAPI(impl appAPI) {
	if AppAPI != nil {
		log.Warn("set AppAPI impl duplicate")
		return
	}
	AppAPI = impl
}
