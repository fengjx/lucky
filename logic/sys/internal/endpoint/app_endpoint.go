package endpoint

import (
	"context"

	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/logic/sys/internal/data/consts"
	"github.com/fengjx/lucky/logic/sys/internal/protocol"
	"github.com/fengjx/lucky/logic/sys/internal/service"
)

var config = configEndpoint{}

type configEndpoint struct {
}

func (e configEndpoint) MakeFetchDataEndpoint() luchen.Endpoint {
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
