package endpoint

import (
	"github.com/fengjx/daox"
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/connom/types"
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/transport/http"
)

type configAdminHandler struct {
}

func (h configAdminHandler) Bind(router *luchen.HTTPServeMux) {
	router.Sub(http.AdminAPI+"/sys/config", func(sub *luchen.HTTPServeMux) {
		sub.Handle("/add", h.add())
		sub.Handle("/update", h.update())
		sub.Handle("/del", h.del())
		sub.Handle("/batch-update", h.batchUpdate())
		sub.Handle("/query", h.query())
	})
}

func (h configAdminHandler) query() *luchen.HTTPTransportServer {
	return http.NewHandler(
		configAdmin.makeQueryEndpoint(),
		luchen.DecodeHTTPJSONRequest[daox.QueryRecord],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h configAdminHandler) add() *luchen.HTTPTransportServer {
	return http.NewHandler(
		configAdmin.makeAddEndpoint(),
		luchen.DecodeHTTPJSONRequest[entity.SysConfig],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h configAdminHandler) update() *luchen.HTTPTransportServer {
	return http.NewHandler(
		configAdmin.makeUpdateEndpoint(),
		luchen.DecodeHTTPJSONRequest[entity.SysConfig],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h configAdminHandler) del() *luchen.HTTPTransportServer {
	return http.NewHandler(
		configAdmin.makeDelEndpoint(),
		luchen.DecodeHTTPJSONRequest[types.DelReq],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h configAdminHandler) batchUpdate() *luchen.HTTPTransportServer {
	return http.NewHandler(
		configAdmin.makeBatchUpdateEndpoint(),
		luchen.DecodeHTTPJSONRequest[types.BatchUpdate],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}
