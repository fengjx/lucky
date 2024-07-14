package endpoint

import (
	"github.com/fengjx/daox"
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/connom/types"
	"github.com/fengjx/lucky/logic/cms/internal/data/entity"
	"github.com/fengjx/lucky/transport/http"
)

type newsAdminHandler struct {
}

func (h newsAdminHandler) Bind(router *luchen.HTTPServeMux) {
	router.Sub(http.AdminAPI+"/cms/news", func(sub *luchen.HTTPServeMux) {
		sub.Handle("/add", h.add())
		sub.Handle("/update", h.update())
		sub.Handle("/del", h.del())
		sub.Handle("/batch-update", h.batchUpdate())
		sub.Handle("/query", h.query())
		sub.Handle("/topics", h.topics())
	})
}

func (h newsAdminHandler) query() *luchen.HTTPTransportServer {
	return http.NewHandler(
		newsAdmin.makeQueryEndpoint(),
		luchen.DecodeHTTPJSONRequest[daox.QueryRecord],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h newsAdminHandler) add() *luchen.HTTPTransportServer {
	return http.NewHandler(
		newsAdmin.makeAddEndpoint(),
		luchen.DecodeHTTPJSONRequest[entity.CmsNews],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h newsAdminHandler) update() *luchen.HTTPTransportServer {
	return http.NewHandler(
		newsAdmin.makeUpdateEndpoint(),
		luchen.DecodeHTTPJSONRequest[entity.CmsNews],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h newsAdminHandler) del() *luchen.HTTPTransportServer {
	return http.NewHandler(
		newsAdmin.makeDelEndpoint(),
		luchen.DecodeHTTPJSONRequest[types.DelReq],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h newsAdminHandler) batchUpdate() *luchen.HTTPTransportServer {
	return http.NewHandler(
		newsAdmin.makeBatchUpdateEndpoint(),
		luchen.DecodeHTTPJSONRequest[types.BatchUpdate],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h newsAdminHandler) topics() *luchen.HTTPTransportServer {
	return http.NewHandler(
		newsAdmin.makeTopicsEndpoint(),
		luchen.NopHTTPRequestDecoder,
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}
