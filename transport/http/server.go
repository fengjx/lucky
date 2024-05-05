package http

import (
	"net/http"
	"sync"

	"github.com/fengjx/luchen"
	"github.com/fengjx/luchen/http/middleware"
	"github.com/fengjx/luchen/http/pprof"

	"github.com/fengjx/lucky/connom/config"
)

var (
	server     *luchen.HTTPServer
	serverOnce sync.Once
)

func GetServer() *luchen.HTTPServer {
	serverOnce.Do(func() {
		serverConfig := config.GetConfig().Server.HTTP
		server = luchen.NewHTTPServer(
			luchen.WithServiceName(serverConfig.ServerName),
			luchen.WithServerAddr(serverConfig.Listen),
		).Use(
			middleware.Recoverer,
			middleware.RequestID,
			middleware.RealIP,
			middleware.CorsHandler(middleware.CorsOptions{
				AllowedOrigins: serverConfig.Cors.AllowOrigins,
				AllowedMethods: []string{
					http.MethodHead,
					http.MethodGet,
					http.MethodPost,
					http.MethodPut,
					http.MethodPatch,
					http.MethodDelete,
				},
				AllowedHeaders: []string{"*"},
				ExposedHeaders: []string{
					"Content-Disposition",
					"Content-Type",
					ResponseHeaderServer,
					ResponseHeaderRefreshToken,
				},
				AllowCredentials: true,
			}),
			commonMiddleware,
			adminMiddleware,
		).Handler(
			pprof.NewHandler().BasicAuth(map[string]string{
				"fengjx": "hello1024",
			}),
		).Static("/static/", "static")
	})
	return server
}
