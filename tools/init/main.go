package main

import (
	"github.com/fengjx/luchen/log"
	"go.uber.org/zap"

	"github.com/fengjx/lucky/logic/sys"
	"github.com/fengjx/lucky/logic/sys/syspub"
)

func main() {
	sys.InitWithTools()
	err := syspub.AppAPI.Init()
	if err != nil {
		log.Error("install err", zap.Error(err))
	}
}
