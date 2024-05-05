package endpoint

import (
	"github.com/fengjx/daox"
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/connom/types"
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/transport/http"
)

type sysDictAdminHandler struct {
}

func (h sysDictAdminHandler) Bind(router *luchen.HTTPServeMux) {
	router.Sub(http.AdminAPI+"/sys/dict", func(sub *luchen.HTTPServeMux) {
		sub.Handle("/add", h.add())
		sub.Handle("/update", h.update())
		sub.Handle("/del", h.del())
		sub.Handle("/batch-update", h.batchUpdate())
		sub.Handle("/query", h.query())
	})
}

func (h sysDictAdminHandler) query() *luchen.HTTPTransportServer {
	return http.NewHandler(
		dictAdmin.makeQueryEndpoint(),
		luchen.DecodeHTTPJSONRequest[daox.QueryRecord],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h sysDictAdminHandler) add() *luchen.HTTPTransportServer {
	return http.NewHandler(
		dictAdmin.makeAddEndpoint(),
		luchen.DecodeHTTPJSONRequest[entity.SysDict],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h sysDictAdminHandler) update() *luchen.HTTPTransportServer {
	return http.NewHandler(
		dictAdmin.makeUpdateEndpoint(),
		luchen.DecodeHTTPJSONRequest[entity.SysDict],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h sysDictAdminHandler) del() *luchen.HTTPTransportServer {
	return http.NewHandler(
		dictAdmin.makeDelEndpoint(),
		luchen.DecodeHTTPJSONRequest[types.DelReq],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h sysDictAdminHandler) batchUpdate() *luchen.HTTPTransportServer {
	return http.NewHandler(
		dictAdmin.makeBatchUpdateEndpoint(),
		luchen.DecodeHTTPJSONRequest[types.BatchUpdate],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}
