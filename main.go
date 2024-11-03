package main

import (
	"time"

	"github.com/fengjx/go-halo/halo"
	"github.com/fengjx/luchen"
	"github.com/fengjx/luchen/log"
	"go.uber.org/zap"

	"github.com/fengjx/lucky/integration"
	"github.com/fengjx/lucky/logic"
	"github.com/fengjx/lucky/transport/http"
)

func main() {
	httpServer := http.GetServer()
	integration.Init()
	logic.Init(httpServer)
	luchen.Start(httpServer)

	if err := halo.Wait(time.Second * 30); err != nil {
		log.Info("server shutdown err", zap.Error(err))
		return
	}
}
