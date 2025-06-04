package user

import (
	"context"

	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/common/types"
	"github.com/fengjx/lucky/logic/sys/internal/protocol"
	"github.com/fengjx/lucky/logic/sys/internal/service"
)

func makeUpdatePwdEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*protocol.UpdateUserPwdReq)
		err = service.UserSvc.UpdatePwd(ctx, req.ID, req.Pwd)
		if err != nil {
			return nil, err
		}
		return types.OKRsp{
			Success: true,
		}, nil
	}
}
