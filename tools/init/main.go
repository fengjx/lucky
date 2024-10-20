package main

import (
	"context"

	"github.com/fengjx/luchen/log"
	"go.uber.org/zap"

	"github.com/fengjx/lucky/logic/sys"
	"github.com/fengjx/lucky/logic/sys/syspub"
)

func main() {
	ctx := context.Background()
	sys.InitWithTools()
	err := syspub.AppAPI.Init(ctx)
	if err != nil {
		log.Error("install err", zap.Error(err))
	}
}
