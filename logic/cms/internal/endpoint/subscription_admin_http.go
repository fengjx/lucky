package endpoint

import (
	"github.com/fengjx/daox"
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/connom/types"
	"github.com/fengjx/lucky/logic/cms/internal/data/entity"
	"github.com/fengjx/lucky/transport/http"
)

type subscriptionAdminHandler struct {
}

func (h subscriptionAdminHandler) Bind(router *luchen.HTTPServeMux) {
	router.Sub(http.AdminAPI+"/cms/subscription", func(sub *luchen.HTTPServeMux) {
		sub.Handle("/add", h.add())
		sub.Handle("/update", h.update())
		sub.Handle("/del", h.del())
		sub.Handle("/batch-update", h.batchUpdate())
		sub.Handle("/query", h.query())
	})
}

func (h subscriptionAdminHandler) query() *luchen.HTTPTransportServer {
	return http.NewHandler(
		subscriptionAdmin.makeQueryEndpoint(),
		luchen.DecodeHTTPJSONRequest[daox.QueryRecord],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h subscriptionAdminHandler) add() *luchen.HTTPTransportServer {
	return http.NewHandler(
		subscriptionAdmin.makeAddEndpoint(),
		luchen.DecodeHTTPJSONRequest[entity.CmsSubscription],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h subscriptionAdminHandler) update() *luchen.HTTPTransportServer {
	return http.NewHandler(
		subscriptionAdmin.makeUpdateEndpoint(),
		luchen.DecodeHTTPJSONRequest[entity.CmsSubscription],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h subscriptionAdminHandler) del() *luchen.HTTPTransportServer {
	return http.NewHandler(
		subscriptionAdmin.makeDelEndpoint(),
		luchen.DecodeHTTPJSONRequest[types.DelReq],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h subscriptionAdminHandler) batchUpdate() *luchen.HTTPTransportServer {
	return http.NewHandler(
		subscriptionAdmin.makeBatchUpdateEndpoint(),
		luchen.DecodeHTTPJSONRequest[types.BatchUpdate],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}
