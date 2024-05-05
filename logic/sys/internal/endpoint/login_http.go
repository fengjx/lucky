package endpoint

import (
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/logic/sys/internal/protocol"
	"github.com/fengjx/lucky/transport/http"
)

type loginHandler struct {
}

func (h *loginHandler) Bind(router *luchen.HTTPServeMux) {
	router.Handle(http.OpenAPI+"/sys/login", h.login())
	router.Handle(http.AdminAPI+"/sys/user/info", h.userInfo())
}

func (h *loginHandler) login() *luchen.HTTPTransportServer {
	return http.NewHandler(
		login.makeLoginEndpoint(),
		luchen.DecodeHTTPJSONRequest[protocol.LoginReq],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h *loginHandler) userInfo() *luchen.HTTPTransportServer {
	return http.NewHandler(
		login.makeUserInfoEndpoint(),
		luchen.NopHTTPRequestDecoder,
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}
