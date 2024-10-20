package syspub

import (
	"github.com/fengjx/luchen/log"
)

var AppAPI appAPI

type appAPI interface {
	Init() error
}

func SetAppAPI(impl appAPI) {
	if AppAPI != nil {
		log.Warn("set AppAPI impl duplicate")
		return
	}
	AppAPI = impl
}
