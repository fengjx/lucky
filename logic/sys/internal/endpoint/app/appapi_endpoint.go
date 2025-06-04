package app

import (
	"context"
	"reflect"

	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/common/types"
	"github.com/fengjx/lucky/logic/sys/internal/data/consts"
	"github.com/fengjx/lucky/logic/sys/internal/protocol"
	"github.com/fengjx/lucky/logic/sys/internal/service"
	"github.com/fengjx/lucky/transport/http"
)

func RegisterAppTTPHandler(hs *luchen.HTTPServer) {
	e := &appEndpoint{}
	hs.Handle(&luchen.EndpointDefine{
		Endpoint: e.makeFetchDataEndpoint(),
		Name:     "App.FetchData",
		Path:     http.OpenAPI + "/app/data",
		ReqType:  reflect.TypeOf(&types.Empty{}),
		RspType:  reflect.TypeOf(&protocol.AppDataResp{}),
	})
}

type appEndpoint struct {
}

func (e *appEndpoint) makeFetchDataEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		appConfig := service.ConfigSvc.ScopeConfig(
			consts.ScopeApp,
		)
		dict := service.DictSvc.GetGroupDict()
		return protocol.AppDataResp{
			Config: appConfig,
			Dict:   dict,
		}, nil
	}
}
