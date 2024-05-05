package service

import (
	"context"

	"github.com/fengjx/daox"
	"github.com/fengjx/daox/sqlbuilder/ql"
	"github.com/fengjx/go-halo/json"
	"github.com/fengjx/go-halo/utils"
	"github.com/fengjx/luchen/log"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"go.uber.org/zap"

	"github.com/fengjx/lucky/connom/types"
	"github.com/fengjx/lucky/integration/db"
	"github.com/fengjx/lucky/logic/common"
	"github.com/fengjx/lucky/logic/sys/internal/dao"
	"github.com/fengjx/lucky/logic/sys/internal/data/dto"
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/internal/data/enum"
	"github.com/fengjx/lucky/logic/sys/internal/data/meta"
)

var MenuBaseSvc = &menuBaseService{}

type menuBaseService struct {
}

// Query 分页查询
func (svc menuBaseService) Query(ctx context.Context, query *daox.QueryRecord) (*types.PageVO[*dto.MenuDTO], error) {
	readDB := dao.SysMenuDao.GetReadDB()
	query.TableName = meta.SysMenuMeta.TableName()
	// 只查根菜单
	query.Conditions = append(query.Conditions, daox.Condition{
		Op:            daox.OpAnd,
		Field:         meta.SysMenuMeta.ParentID,
		ConditionType: daox.ConditionTypeEq,
		Vals:          []any{0},
	})
	query.OrderFields = append(query.OrderFields, daox.OrderField{Field: meta.SysMenuMeta.SortNo, OrderType: daox.OrderTypeAsc})
	list, page, err := daox.Find[*entity.SysMenu](ctx, readDB, *query)
	if err != nil {
		log.ErrorCtx(ctx, "page query sys_menu err", zap.Any("query", json.ToJsonDelay(query)), zap.Error(err))
		return nil, err
	}
	var parentIDs []int64
	dtos := lo.Map[*entity.SysMenu, *dto.MenuDTO](list, func(item *entity.SysMenu, index int) *dto.MenuDTO {
		parentIDs = append(parentIDs, item.ID)
		return dto.BuildMenuDTO(item)
	})
	var status []enum.MenuStatus
	val := common.GetConditionVal(meta.SysMenuMeta.Status, query.Conditions)
	if len(val) == 0 {
		status = []enum.MenuStatus{
			enum.MenuStatusNormal,
			enum.MenuStatusDisable,
		}
	} else {
		status = []enum.MenuStatus{
			enum.MenuStatus(utils.ToString(val[0])),
		}
	}

	menuMap, err := svc.recursiveChildren(ctx, status, parentIDs...)
	if err != nil {
		return nil, err
	}
	for _, item := range dtos {
		item.Children = menuMap[item.ID]
	}
	pageVO := &types.PageVO[*dto.MenuDTO]{
		List:    dtos,
		Offset:  page.Offset,
		Limit:   page.Limit,
		Count:   page.Count,
		HasNext: page.HasNext,
	}
	return pageVO, nil
}

// Add 新增记录
func (svc menuBaseService) Add(ctx context.Context, model *entity.SysMenu) (int64, error) {
	return dao.SysMenuDao.SaveContext(ctx, model,
		meta.SysMenuMeta.Ctime,
		meta.SysMenuMeta.Utime,
		meta.SysMenuMeta.IsSys,
	)
}

// Update 更新记录
func (svc menuBaseService) Update(ctx context.Context, model *entity.SysMenu) (bool, error) {
	return dao.SysMenuDao.UpdateContext(ctx, model,
		meta.SysMenuMeta.PrimaryKey(),
		meta.SysMenuMeta.Ctime,
		meta.SysMenuMeta.Utime,
		meta.SysMenuMeta.IsSys,
	)
}

// BatchUpdate 批量更新
func (svc menuBaseService) BatchUpdate(ctx context.Context, param *types.BatchUpdate) (bool, error) {
	for _, row := range param.Rows {
		var id any
		attr := map[string]any{}
		for k, v := range row {
			if k == meta.SysMenuMeta.PrimaryKey() {
				id = v
				continue
			}
			attr[k] = v
		}
		err := db.GetDefaultTxManager().ExecTx(ctx, func(txCtx context.Context, tx *sqlx.Tx) error {
			_, err := dao.SysMenuDao.UpdateFieldTxContext(txCtx, tx, id, attr)
			return err
		})
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// DeleteByIDs 批量更新
func (svc menuBaseService) DeleteByIDs(ctx context.Context, ids []int64) error {
	l := log.GetLogger(ctx).With(zap.Any("ids", ids))
	_, err := dao.SysMenuDao.DeleteByCondContext(ctx, ql.C().And(
		meta.SysMenuMeta.IdIn(ids...),
		meta.SysMenuMeta.IsSysNotEQ(1),
	))
	if err != nil {
		l.Error("delete sys_menu err", zap.Error(err))
		return err
	}
	l.Info("delete sys_menu success")
	return nil
}

// TreeList 查询菜单 tree
func (svc menuBaseService) TreeList(ctx context.Context, status []enum.MenuStatus) ([]*dto.MenuDTO, error) {
	rootList, err := dao.SysMenuDao.ListChildren(ctx, status, 0)
	if err != nil {
		log.ErrorCtx(ctx, "query root menu list err", zap.Error(err))
		return nil, err
	}
	if len(rootList) == 0 {
		return nil, nil
	}
	root := rootList[0]
	var parentIDs []int64
	dtos := lo.Map[*entity.SysMenu, *dto.MenuDTO](root, func(item *entity.SysMenu, index int) *dto.MenuDTO {
		parentIDs = append(parentIDs, item.ID)
		return dto.BuildMenuDTO(item)
	})
	menuMap, err := svc.recursiveChildren(ctx, status, parentIDs...)
	if err != nil {
		log.ErrorCtx(ctx, "recursive query menu children err", zap.Error(err))
		return nil, err
	}
	for _, item := range dtos {
		item.Children = menuMap[item.ID]
	}
	return dtos, nil
}

// 递归获取子菜单
func (svc menuBaseService) recursiveChildren(ctx context.Context, status []enum.MenuStatus, parentIDs ...int64) (map[int64][]*dto.MenuDTO, error) {
	if len(parentIDs) == 0 {
		return nil, nil
	}
	children, err := dao.SysMenuDao.ListChildren(ctx, status, parentIDs...)
	if err != nil {
		return nil, err
	}
	res := map[int64][]*dto.MenuDTO{}
	for parentID, menus := range children {
		if len(menus) == 0 {
			continue
		}
		var pids []int64
		var dtos []*dto.MenuDTO
		for _, menu := range menus {
			pids = append(pids, menu.ID)
			dtos = append(dtos, dto.BuildMenuDTO(menu))
		}
		rChildren, err := svc.recursiveChildren(ctx, status, pids...)
		if err != nil {
			return nil, err
		}
		for _, menuDTO := range dtos {
			menuDTO.Children = rChildren[menuDTO.ID]
		}
		res[parentID] = dtos
	}
	return res, nil
}
