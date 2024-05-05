package endpoint

import (
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/transport/http"
)

type configHandler struct {
}

func (h configHandler) Bind(router *luchen.HTTPServeMux) {
	router.Handle(http.OpenAPI+"/app/data", h.fetchData())
}

func (h configHandler) fetchData() *luchen.HTTPTransportServer {
	return http.NewHandler(
		config.MakeFetchDataEndpoint(),
		luchen.NopHTTPRequestDecoder,
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}
