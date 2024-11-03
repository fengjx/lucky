package endpoint

import (
	"context"
	"strconv"

	"github.com/fengjx/daox"
	"github.com/fengjx/go-halo/errs"
	"github.com/fengjx/go-halo/utils"
	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/connom/types"
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/internal/service"
)

var configAdmin = configAdminEndpoint{}

type configAdminEndpoint struct {
}

func (e configAdminEndpoint) makeAddEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.SysConfig)
		id, err := service.ConfigBaseSvc.Add(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "add sys_config err")
		}
		response = types.AddRsp{
			ID: id,
		}
		return
	}
}

func (e configAdminEndpoint) makeUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.SysConfig)
		ok, err := service.ConfigBaseSvc.Update(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "update sys_config err")
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e configAdminEndpoint) makeDelEndpoint() luchen.Endpoint {
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
		err = service.ConfigBaseSvc.DeleteByIDs(ctx, ids)
		if err != nil {
			return nil, errs.Wrap(err, "delete sys_config err")
		}
		return
	}
}

func (e configAdminEndpoint) makeBatchUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*types.BatchUpdate)
		ok, err := service.ConfigBaseSvc.BatchUpdate(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "batch update sys_config err")
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e configAdminEndpoint) makeQueryEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		query := request.(*daox.QueryRecord)
		pageVO, err := service.ConfigBaseSvc.Query(ctx, query)
		if err != nil {
			return nil, errs.Wrap(err, "page query sys_config err")
		}
		return pageVO.ToAmisResp(), nil
	}
}
