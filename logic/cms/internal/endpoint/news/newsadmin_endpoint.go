package news

import (
	"context"
	"reflect"
	"strconv"

	"github.com/fengjx/daox"
	"github.com/fengjx/go-halo/errs"
	"github.com/fengjx/go-halo/utils"
	"github.com/fengjx/luchen"
	"github.com/fengjx/lucky/common/types"
	"github.com/fengjx/lucky/logic/cms/internal/data/entity"
	"github.com/fengjx/lucky/logic/cms/internal/service"
)

func RegisterNewsAdminTTPHandler(hs *luchen.HTTPServer) {
	e := &newsAdminEndpoint{}
	hs.Handle(&luchen.EndpointDefine{
		Endpoint: e.makeAddEndpoint(),
		Name:     "NewsAdmin.Add",
		Path:     "/admin/cms/news/add",
		ReqType:  reflect.TypeOf(&entity.CmsNews{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "NewsAdmin.Update",
		Path:     "/admin/cms/news/update",
		ReqType:  reflect.TypeOf(&entity.CmsNews{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeUpdateEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "NewsAdmin.Del",
		Path:     "/admin/cms/news/del",
		ReqType:  reflect.TypeOf(&types.DelReq{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeDelEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "NewsAdmin.BatchUpdate",
		Path:     "/admin/cms/news/batch-update",
		ReqType:  reflect.TypeOf(&types.BatchUpdate{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeBatchUpdateEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "NewsAdmin.Query",
		Path:     "/admin/cms/news/query",
		ReqType:  reflect.TypeOf(&daox.QueryRecord{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeQueryEndpoint(),
	})
}

type newsAdminEndpoint struct {
}

func (e newsAdminEndpoint) makeAddEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.CmsNews)
		id, err := service.NewsBaseSvc.Add(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "add cms_news err")
		}
		response = types.AddRsp{
			ID: id,
		}
		return
	}
}

func (e newsAdminEndpoint) makeUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.CmsNews)
		ok, err := service.NewsBaseSvc.Update(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "update cms_news err")
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e newsAdminEndpoint) makeDelEndpoint() luchen.Endpoint {
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
		err = service.NewsBaseSvc.DeleteByIDs(ctx, ids)
		if err != nil {
			return nil, errs.Wrap(err, "delete cms_news err")
		}
		return
	}
}

func (e newsAdminEndpoint) makeBatchUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*types.BatchUpdate)
		ok, err := service.NewsBaseSvc.BatchUpdate(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "batch update cms_news err")
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e newsAdminEndpoint) makeQueryEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		query := request.(*daox.QueryRecord)
		pageVO, err := service.NewsBaseSvc.Query(ctx, query)
		if err != nil {
			return nil, errs.Wrap(err, "page query cms_news err")
		}
		return pageVO.ToAmisResp(), nil
	}
}

func (e newsAdminEndpoint) makeTopicsEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		topics, err := service.TopicSvc.ListAll(ctx)
		if err != nil {
			return nil, errs.Wrap(err, "find all topic err")
		}
		var res []types.Option
		for _, topic := range topics {
			res = append(res, types.Option{
				Label: topic.Name,
				Value: topic.Code,
			})
		}
		return res, nil
	}
}
