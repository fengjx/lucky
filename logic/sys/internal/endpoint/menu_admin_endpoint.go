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
	"github.com/fengjx/lucky/logic/sys/internal/data/dto"
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/internal/data/enum"
	"github.com/fengjx/lucky/logic/sys/internal/protocol"
	"github.com/fengjx/lucky/logic/sys/internal/service"
)

var menuAdmin = sysMenuAdminEndpoint{}

type sysMenuAdminEndpoint struct {
}

func (e sysMenuAdminEndpoint) makeAddEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.SysMenu)
		id, err := service.MenuBaseSvc.Add(ctx, param)
		if err != nil {
			log.ErrorCtx(ctx, "add sys_menu err", zap.Any("param", param), zap.Error(err))
			return nil, err
		}
		response = types.AddRsp{
			ID: id,
		}
		return
	}
}

func (e sysMenuAdminEndpoint) makeUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*entity.SysMenu)
		ok, err := service.MenuBaseSvc.Update(ctx, param)
		if err != nil {
			log.ErrorCtx(ctx, "update sys_menu err", zap.Any("param", param), zap.Error(err))
			return nil, err
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e sysMenuAdminEndpoint) makeDelEndpoint() luchen.Endpoint {
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

func (e sysMenuAdminEndpoint) makeBatchUpdateEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(*types.BatchUpdate)
		ok, err := service.MenuBaseSvc.BatchUpdate(ctx, param)
		if err != nil {
			log.ErrorCtx(ctx, "batch update sys_menu err", zap.Any("param", param), zap.Error(err))
			return nil, err
		}
		response = types.OKRsp{
			Success: ok,
		}
		return
	}
}

func (e sysMenuAdminEndpoint) makeQueryEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		query := request.(*daox.QueryRecord)
		pageVO, err := service.MenuBaseSvc.Query(ctx, query)
		if err != nil {
			log.ErrorCtx(ctx, "page query sys_menu err", zap.Error(err))
			return nil, err
		}
		return pageVO.ToAmisResp(), nil
	}
}

func (e sysMenuAdminEndpoint) buildOption(menus []*dto.MenuDTO) []*types.Option {
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

func (e sysMenuAdminEndpoint) makeOptionsEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		treeList, err := service.MenuBaseSvc.TreeList(ctx, []enum.MenuStatus{
			enum.MenuStatusNormal,
			enum.MenuStatusDisable,
		})
		if err != nil {
			log.ErrorCtx(ctx, "recursive query menus err", zap.Error(err))
			return nil, err
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

func (e sysMenuAdminEndpoint) buildMenu(menus []*dto.MenuDTO) []*protocol.Menu {
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

func (e sysMenuAdminEndpoint) makeFetchEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		treeList, err := service.MenuBaseSvc.TreeList(ctx, []enum.MenuStatus{
			enum.MenuStatusNormal,
		})
		if err != nil {
			log.ErrorCtx(ctx, "recursive query menus err", zap.Error(err))
			return nil, err
		}
		pages := e.buildMenu(treeList)
		rsp := protocol.App{
			Pages: pages,
		}
		return rsp, nil
	}
}
