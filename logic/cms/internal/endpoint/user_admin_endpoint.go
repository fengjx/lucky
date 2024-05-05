package endpoint

import (
	"context"
	"strconv"

	"github.com/fengjx/daox"
	"github.com/fengjx/go-halo/utils"
	"github.com/fengjx/luchen"
	"github.com/fengjx/luchen/log"
	"go.uber.org/zap"

	"github.com/fengjx/lucky/connom/types"
	"github.com/fengjx/lucky/logic/cms/internal/data/entity"
	"github.com/fengjx/lucky/logic/cms/internal/service"
)

var userAdmin = userAdminEndpoint{}

type userAdminEndpoint struct {
}

func (e userAdminEndpoint) makeAddEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.CmsUser)
		id, err := service.UserBaseSvc.Add(ctx, param)
		if err != nil {
			log.ErrorCtx(ctx, "add cms_user err", zap.Any("param", param), zap.Error(err))
			return nil, err
		}
		response = types.AddRsp{
			ID: id,
		}
		return
	}
}

func (e userAdminEndpoint) makeUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.CmsUser)
		ok, err := service.UserBaseSvc.Update(ctx, param)
		if err != nil {
			log.ErrorCtx(ctx, "update cms_user err", zap.Any("param", param), zap.Error(err))
			return nil, err
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e userAdminEndpoint) makeDelEndpoint() luchen.Endpoint {
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
			return nil, err
		}
		return
	}
}

func (e userAdminEndpoint) makeBatchUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*types.BatchUpdate)
		ok, err := service.UserBaseSvc.BatchUpdate(ctx, param)
		if err != nil {
			log.ErrorCtx(ctx, "batch update cms_user err", zap.Any("param", param), zap.Error(err))
			return nil, err
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e userAdminEndpoint) makeQueryEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		query := request.(*daox.QueryRecord)
		pageVO, err := service.UserBaseSvc.Query(ctx, query)
		if err != nil {
			log.ErrorCtx(ctx, "page query cms_user err", zap.Error(err))
			return nil, err
		}
		return pageVO.ToAmisResp(), nil
	}
}
