package endpoint

import (
	"github.com/fengjx/daox"
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/connom/types"
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/transport/http"
)

type sysMenuAdminHandler struct {
}

func (h sysMenuAdminHandler) Bind(router *luchen.HTTPServeMux) {
	router.Sub(http.AdminAPI+"/sys/menu", func(sub *luchen.HTTPServeMux) {
		sub.Handle("/add", h.add())
		sub.Handle("/update", h.update())
		sub.Handle("/del", h.del())
		sub.Handle("/batch-update", h.batchUpdate())
		sub.Handle("/query", h.query())
		sub.Handle("/options", h.options())
		sub.Handle("/fetch", h.fetch())
	})
}

func (h sysMenuAdminHandler) query() *luchen.HTTPTransportServer {
	return http.NewHandler(
		menuAdmin.makeQueryEndpoint(),
		luchen.DecodeHTTPJSONRequest[daox.QueryRecord],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h sysMenuAdminHandler) add() *luchen.HTTPTransportServer {
	return http.NewHandler(
		menuAdmin.makeAddEndpoint(),
		luchen.DecodeHTTPJSONRequest[entity.SysMenu],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h sysMenuAdminHandler) update() *luchen.HTTPTransportServer {
	return http.NewHandler(
		menuAdmin.makeUpdateEndpoint(),
		luchen.DecodeHTTPJSONRequest[entity.SysMenu],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h sysMenuAdminHandler) del() *luchen.HTTPTransportServer {
	return http.NewHandler(
		menuAdmin.makeDelEndpoint(),
		luchen.DecodeHTTPJSONRequest[types.DelReq],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h sysMenuAdminHandler) batchUpdate() *luchen.HTTPTransportServer {
	return http.NewHandler(
		menuAdmin.makeBatchUpdateEndpoint(),
		luchen.DecodeHTTPJSONRequest[types.BatchUpdate],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h sysMenuAdminHandler) options() *luchen.HTTPTransportServer {
	return http.NewHandler(
		menuAdmin.makeOptionsEndpoint(),
		luchen.NopHTTPRequestDecoder,
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h sysMenuAdminHandler) fetch() *luchen.HTTPTransportServer {
	return http.NewHandler(
		menuAdmin.makeFetchEndpoint(),
		luchen.NopHTTPRequestDecoder,
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}
