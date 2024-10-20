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

var dictAdmin = sysDictAdminEndpoint{}

type sysDictAdminEndpoint struct {
}

func (e sysDictAdminEndpoint) makeAddEndpoint() luchen.Endpoint {
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

func (e sysDictAdminEndpoint) makeUpdateEndpoint() luchen.Endpoint {
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

func (e sysDictAdminEndpoint) makeDelEndpoint() luchen.Endpoint {
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

func (e sysDictAdminEndpoint) makeBatchUpdateEndpoint() luchen.Endpoint {
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

func (e sysDictAdminEndpoint) makeQueryEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		query := request.(*daox.QueryRecord)
		pageVO, err := service.DictBaseSvc.Query(ctx, query)
		if err != nil {
			return nil, errs.Wrap(err, "page query sys_dict err")
		}
		return pageVO.ToAmisResp(), nil
	}
}
