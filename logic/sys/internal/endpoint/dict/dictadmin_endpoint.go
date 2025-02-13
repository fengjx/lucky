package dict

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
	"github.com/fengjx/lucky/logic/sys/internal/service"
)

// RegisterDictAdminTTPHandler 注册 HTTP 路由
func RegisterDictAdminTTPHandler(hs *luchen.HTTPServer) {
	e := &dictAdminEndpoint{}
	hs.Handle(&luchen.EndpointDefine{
		Name:     "DictAdmin.Add",
		Path:     "/admin/sys/dict/add",
		ReqType:  reflect.TypeOf(&entity.SysDict{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeAddEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "DictAdmin.Update",
		Path:     "/admin/sys/dict/update",
		ReqType:  reflect.TypeOf(&entity.SysDict{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeUpdateEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "DictAdmin.Del",
		Path:     "/admin/sys/dict/del",
		ReqType:  reflect.TypeOf(&types.DelReq{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeDelEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "DictAdmin.BatchUpdate",
		Path:     "/admin/sys/dict/batch-update",
		ReqType:  reflect.TypeOf(&types.BatchUpdate{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeBatchUpdateEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "DictAdmin.Query",
		Path:     "/admin/sys/dict/query",
		ReqType:  reflect.TypeOf(&daox.QueryRecord{}),
		RspType:  reflect.TypeOf(&types.AmisPageResp[*entity.SysDict]{}),
		Endpoint: e.makeQueryEndpoint(),
	})
}

type dictAdminEndpoint struct {
}

func (e *dictAdminEndpoint) makeAddEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.SysDict)
		id, err := service.DictBaseSvc.Add(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "add sys_dict err")
		}
		response = types.AddRsp{
			ID: id,
		}
		return
	}
}

func (e *dictAdminEndpoint) makeUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.SysDict)
		ok, err := service.DictBaseSvc.Update(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "update sys_dict err")
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e *dictAdminEndpoint) makeDelEndpoint() luchen.Endpoint {
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
		err = service.DictBaseSvc.DeleteByIDs(ctx, ids)
		if err != nil {
			return nil, errs.Wrap(err, "delete sys_dict err")
		}
		return
	}
}

func (e *dictAdminEndpoint) makeBatchUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*types.BatchUpdate)
		ok, err := service.DictBaseSvc.BatchUpdate(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "batch update sys_dict err")
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e *dictAdminEndpoint) makeQueryEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		query := request.(*daox.QueryRecord)
		pageVO, err := service.DictBaseSvc.Query(ctx, query)
		if err != nil {
			return nil, errs.Wrap(err, "page query sys_dict err")
		}
		return pageVO.ToAmisResp(), nil
	}
}
