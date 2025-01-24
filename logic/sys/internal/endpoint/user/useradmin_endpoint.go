package user

import (
	"context"
	"reflect"
	"strconv"

	"github.com/fengjx/daox"
	"github.com/fengjx/go-halo/errs"
	"github.com/fengjx/go-halo/utils"
	"github.com/fengjx/luchen"
	"github.com/fengjx/lucky/common/types"
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/internal/protocol"
	"github.com/fengjx/lucky/logic/sys/internal/service"
)

func RegisterUserAdminEndpoint(hs *luchen.HTTPServer) {
	e := &userAdminEndpoint{}
	hs.Handle(&luchen.EndpointDefine{
		Name:     "UserAdmin.Add",
		Path:     "/admin/sys/user/add",
		ReqType:  reflect.TypeOf(&entity.SysUser{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeAddEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "UserAdmin.Update",
		Path:     "/admin/sys/user/update",
		ReqType:  reflect.TypeOf(&entity.SysUser{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeUpdateEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "UserAdmin.Del",
		Path:     "/admin/sys/user/del",
		ReqType:  reflect.TypeOf(&types.DelReq{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeDelEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "UserAdmin.BatchUpdate",
		Path:     "/admin/sys/user/batch-update",
		ReqType:  reflect.TypeOf(&types.BatchUpdate{}),
		RspType:  reflect.TypeOf(&types.AmisPageResp[*entity.SysUser]{}),
		Endpoint: e.makeBatchUpdateEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "UserAdmin.Query",
		Path:     "/admin/sys/user/query",
		ReqType:  reflect.TypeOf(&daox.QueryRecord{}),
		RspType:  reflect.TypeOf(&types.AmisPageResp[*entity.SysUser]{}),
		Endpoint: e.makeQueryEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "UserAdmin.UpdatePwd",
		Path:     "/admin/sys/user/update-pwd",
		ReqType:  reflect.TypeOf(&protocol.UpdateUserPwdReq{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeUpdatePwdEndpoint(),
	})
}

type userAdminEndpoint struct {
}

func (e *userAdminEndpoint) makeAddEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.SysUser)
		id, err := service.UserBaseSvc.Add(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "add user err")
		}
		response = types.AddRsp{
			ID: id,
		}
		return
	}
}

func (e *userAdminEndpoint) makeUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.SysUser)
		ok, err := service.UserBaseSvc.Update(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "update user err")
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e *userAdminEndpoint) makeDelEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*types.DelReq)
		res := types.OKRsp{Success: true}
		if param.IDs == "" {
			return res, nil
		}
		ids := utils.SplitToSlice[int64](param.IDs, ",", func(item string) int64 {
			i, _ := strconv.ParseInt(item, 10, 64)
			return i
		})
		err = service.UserBaseSvc.DeleteByIDs(ctx, ids)
		if err != nil {
			return nil, errs.Wrap(err, "delete user err")
		}
		return
	}
}

func (e *userAdminEndpoint) makeBatchUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*types.BatchUpdate)
		ok, err := service.UserBaseSvc.BatchUpdate(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "batch update user err")
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e *userAdminEndpoint) makeQueryEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		query := request.(*daox.QueryRecord)
		pageVO, err := service.UserBaseSvc.Query(ctx, query)
		if err != nil {
			return nil, errs.Wrap(err, "page query user err")
		}
		return pageVO.ToAmisResp(), nil
	}
}

func (e *userAdminEndpoint) makeUpdatePwdEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*protocol.UpdateUserPwdReq)
		err = service.UserBaseSvc.UpdatePwd(ctx, req.ID, req.Pwd)
		if err != nil {
			return nil, err
		}
		return types.OKRsp{
			Success: true,
		}, nil
	}
}
