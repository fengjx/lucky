package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fengjx/luchen"
	"github.com/fengjx/luchen/log"

	"github.com/fengjx/lucky/logic"
	"github.com/fengjx/lucky/transport/http"
)

func main() {
	ctx := context.Background()
	httpServer := http.GetServer()
	logic.Init(ctx, httpServer)
	luchen.Start(httpServer)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-quit
	log.Info("server shutdown")
	luchen.Stop()
}
