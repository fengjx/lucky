package menu

import (
	"context"
	"reflect"
	"strconv"

	"github.com/fengjx/daox"
	"github.com/fengjx/go-halo/errs"
	"github.com/fengjx/go-halo/utils"
	"github.com/fengjx/luchen"
	"github.com/fengjx/lucky/common/types"
	"github.com/fengjx/lucky/logic/sys/internal/data/dto"
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/internal/data/enum"
	"github.com/fengjx/lucky/logic/sys/internal/protocol"
	"github.com/fengjx/lucky/logic/sys/internal/service"
)

func RegisterMenuAdminTTPHandler(hs *luchen.HTTPServer) {
	e := &menuAdminEndpoint{}
	hs.Handle(&luchen.EndpointDefine{
		Name:     "MenuAdmin.Add",
		Path:     "/admin/sys/menu/add",
		ReqType:  reflect.TypeOf(&entity.SysMenu{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeAddEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "MenuAdmin.Update",
		Path:     "/admin/sys/menu/update",
		ReqType:  reflect.TypeOf(&entity.SysMenu{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeUpdateEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "MenuAdmin.Del",
		Path:     "/admin/sys/menu/del",
		ReqType:  reflect.TypeOf(&types.DelReq{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeDelEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "MenuAdmin.BatchUpdate",
		Path:     "/admin/sys/menu/batch-update",
		ReqType:  reflect.TypeOf(&types.BatchUpdate{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeBatchUpdateEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "MenuAdmin.Query",
		Path:     "/admin/sys/menu/query",
		ReqType:  reflect.TypeOf(&daox.QueryRecord{}),
		RspType:  reflect.TypeOf(&types.AddRsp{}),
		Endpoint: e.makeQueryEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "MenuAdmin.Options",
		Path:     "/admin/sys/menu/options",
		ReqType:  reflect.TypeOf(&types.Empty{}),
		RspType:  reflect.TypeOf(&types.SelectResp{}),
		Endpoint: e.makeOptionsEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "MenuAdmin.Fetch",
		Path:     "/admin/sys/menu/fetch",
		ReqType:  reflect.TypeOf(&types.Empty{}),
		RspType:  reflect.TypeOf(&protocol.App{}),
		Endpoint: e.makeFetchEndpoint(),
	})

}

type menuAdminEndpoint struct {
}

func (e *menuAdminEndpoint) makeAddEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.SysMenu)
		id, err := service.MenuBaseSvc.Add(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "add sys_menu err")
		}
		response = types.AddRsp{
			ID: id,
		}
		return
	}
}

func (e *menuAdminEndpoint) makeUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.SysMenu)
		ok, err := service.MenuBaseSvc.Update(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "update sys_menu err")
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e *menuAdminEndpoint) makeDelEndpoint() luchen.Endpoint {
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
		err = service.MenuBaseSvc.DeleteByIDs(ctx, ids)
		if err != nil {
			return nil, err
		}
		return
	}
}

func (e *menuAdminEndpoint) makeBatchUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*types.BatchUpdate)
		ok, err := service.MenuBaseSvc.BatchUpdate(ctx, param)
		if err != nil {
			return nil, errs.Wrap(err, "batch update sys_menu err")
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e *menuAdminEndpoint) makeQueryEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		query := request.(*daox.QueryRecord)
		pageVO, err := service.MenuBaseSvc.Query(ctx, query)
		if err != nil {
			return nil, errs.Wrap(err, "page query sys_menu err")
		}
		return pageVO.ToAmisResp(), nil
	}
}

func (e *menuAdminEndpoint) buildOption(menus []*dto.MenuDTO) []*types.Option {
	if len(menus) == 0 {
		return nil
	}
	var options []*types.Option
	for _, menu := range menus {
		opt := &types.Option{
			Label:    menu.Name,
			Value:    menu.ID,
			Children: e.buildOption(menu.Children),
		}
		options = append(options, opt)
	}
	return options
}

func (e *menuAdminEndpoint) makeOptionsEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		treeList, err := service.MenuBaseSvc.TreeList(ctx, []enum.MenuStatus{
			enum.MenuStatusNormal,
			enum.MenuStatusDisable,
		})
		if err != nil {
			return nil, errs.Wrap(err, "recursive query menus err")
		}
		options := []*types.Option{
			{
				Label:    "根目录",
				Value:    0,
				Children: e.buildOption(treeList),
			},
		}
		rsp := types.SelectResp{
			Options: options,
		}
		return rsp, nil
	}
}

func (e *menuAdminEndpoint) buildMenu(menus []*dto.MenuDTO) []*protocol.Menu {
	if len(menus) == 0 {
		return nil
	}
	var options []*protocol.Menu
	for _, menu := range menus {
		opt := &protocol.Menu{
			Label:     menu.Name,
			Icon:      menu.Icon,
			URL:       menu.Path,
			Redirect:  menu.Redirect,
			SchemaAPI: menu.SchemaAPI,
			Visible:   menu.Visible == 1,
			Children:  e.buildMenu(menu.Children),
		}
		options = append(options, opt)
	}
	return options
}

func (e *menuAdminEndpoint) makeFetchEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		treeList, err := service.MenuBaseSvc.TreeList(ctx, []enum.MenuStatus{
			enum.MenuStatusNormal,
		})
		if err != nil {
			return nil, errs.Wrap(err, "recursive query menus err")
		}
		pages := e.buildMenu(treeList)
		rsp := protocol.App{
			Pages: pages,
		}
		return rsp, nil
	}
}
