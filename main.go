package main

import (
	"time"

	"github.com/fengjx/go-halo/halo"
	"github.com/fengjx/luchen"
	"github.com/fengjx/luchen/log"
	"go.uber.org/zap"

	"github.com/fengjx/lucky/logic"
	"github.com/fengjx/lucky/logic/admin"
	"github.com/fengjx/lucky/middleware"
	"github.com/fengjx/lucky/transport/http"
)

func main() {
	// 注册全局拦截器
	luchen.UseGlobalHTTPMiddleware(
		luchen.LogMiddleware,
		middleware.AccessMiddleware,
	)
	hs := http.GetServer()
	hs.Mux().GroupHandler(admin.APIPrefix, admin.Router)
	logic.Init(hs)
	luchen.Start(hs)

	if err := halo.Wait(time.Second * 30); err != nil {
		log.Info("server shutdown err", zap.Error(err))
		return
	}
}
