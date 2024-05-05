package endpoint

import (
	"github.com/fengjx/daox"
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/connom/types"
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/internal/protocol"
	"github.com/fengjx/lucky/transport/http"
)

type userAdminHandler struct {
}

func (h userAdminHandler) Bind(router *luchen.HTTPServeMux) {
	router.Sub(http.AdminAPI+"/sys/user", func(sub *luchen.HTTPServeMux) {
		sub.Handle("/add", h.add())
		sub.Handle("/update", h.update())
		sub.Handle("/update-pwd", h.updatePwd())
		sub.Handle("/del", h.del())
		sub.Handle("/batch-update", h.batchUpdate())
		sub.Handle("/query", h.query())
	})
}

func (h userAdminHandler) query() *luchen.HTTPTransportServer {
	return http.NewHandler(
		userAdmin.makeQueryEndpoint(),
		luchen.DecodeHTTPJSONRequest[daox.QueryRecord],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h userAdminHandler) add() *luchen.HTTPTransportServer {
	return http.NewHandler(
		userAdmin.makeAddEndpoint(),
		luchen.DecodeHTTPJSONRequest[entity.SysUser],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h userAdminHandler) update() *luchen.HTTPTransportServer {
	return http.NewHandler(
		userAdmin.makeUpdateEndpoint(),
		luchen.DecodeHTTPJSONRequest[entity.SysUser],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h userAdminHandler) del() *luchen.HTTPTransportServer {
	return http.NewHandler(
		userAdmin.makeDelEndpoint(),
		luchen.DecodeHTTPJSONRequest[types.DelReq],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h userAdminHandler) batchUpdate() *luchen.HTTPTransportServer {
	return http.NewHandler(
		userAdmin.makeBatchUpdateEndpoint(),
		luchen.DecodeHTTPJSONRequest[types.BatchUpdate],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}

func (h userAdminHandler) updatePwd() *luchen.HTTPTransportServer {
	return http.NewHandler(
		userAdmin.makeUpdatePwdEndpoint(),
		luchen.DecodeHTTPJSONRequest[protocol.UpdateUserPwdReq],
		luchen.EncodeHTTPJSONResponse(http.ResponseWrapper),
	)
}
