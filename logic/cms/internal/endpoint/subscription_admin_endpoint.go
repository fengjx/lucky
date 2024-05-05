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

var subscriptionAdmin = subscriptionAdminEndpoint{}

type subscriptionAdminEndpoint struct {
}

func (e subscriptionAdminEndpoint) makeAddEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.CmsSubscription)
		id, err := service.SubscriptionBaseSvc.Add(ctx, param)
		if err != nil {
			log.ErrorCtx(ctx, "add cms_subscription err", zap.Any("param", param), zap.Error(err))
			return nil, err
		}
		response = types.AddRsp{
			ID: id,
		}
		return
	}
}

func (e subscriptionAdminEndpoint) makeUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.CmsSubscription)
		ok, err := service.SubscriptionBaseSvc.Update(ctx, param)
		if err != nil {
			log.ErrorCtx(ctx, "update cms_subscription err", zap.Any("param", param), zap.Error(err))
			return nil, err
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e subscriptionAdminEndpoint) makeDelEndpoint() luchen.Endpoint {
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
		err = service.SubscriptionBaseSvc.DeleteByIDs(ctx, ids)
		if err != nil {
			return nil, err
		}
		return
	}
}

func (e subscriptionAdminEndpoint) makeBatchUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*types.BatchUpdate)
		ok, err := service.SubscriptionBaseSvc.BatchUpdate(ctx, param)
		if err != nil {
			log.ErrorCtx(ctx, "batch update cms_subscription err", zap.Any("param", param), zap.Error(err))
			return nil, err
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e subscriptionAdminEndpoint) makeQueryEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		query := request.(*daox.QueryRecord)
		pageVO, err := service.SubscriptionBaseSvc.Query(ctx, query)
		if err != nil {
			log.ErrorCtx(ctx, "page query cms_subscription err", zap.Error(err))
			return nil, err
		}
		return pageVO.ToAmisResp(), nil
	}
}
