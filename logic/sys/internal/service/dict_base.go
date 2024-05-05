package service

import (
	"context"

	"github.com/fengjx/daox"
	"github.com/fengjx/daox/sqlbuilder/ql"
	"github.com/fengjx/go-halo/json"
	"github.com/fengjx/luchen/log"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/fengjx/lucky/connom/types"
	"github.com/fengjx/lucky/integration/db"
	"github.com/fengjx/lucky/logic/sys/internal/dao"
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/internal/data/meta"
)

var DictBaseSvc = &dictBaseService{}

type dictBaseService struct {
}

// Query 分页查询
func (svc dictBaseService) Query(ctx context.Context, query *daox.QueryRecord) (*types.PageVO[entity.SysDict], error) {
	readDB := dao.SysDictDao.GetReadDB()
	query.TableName = meta.SysDictMeta.TableName()
	list, page, err := daox.Find[entity.SysDict](ctx, readDB, *query)
	if err != nil {
		log.ErrorCtx(ctx, "page query sys_dict err", zap.Any("query", json.ToJsonDelay(query)), zap.Error(err))
		return nil, err
	}
	pageVO := &types.PageVO[entity.SysDict]{
		List:    list,
		Offset:  page.Offset,
		Limit:   page.Limit,
		Count:   page.Count,
		HasNext: page.HasNext,
	}
	return pageVO, nil
}

// Add 新增记录
func (svc dictBaseService) Add(ctx context.Context, model *entity.SysDict) (int64, error) {
	return dao.SysDictDao.SaveContext(ctx, model,
		meta.SysDictMeta.Ctime,
		meta.SysDictMeta.Utime,
	)
}

// Update 更新记录
func (svc dictBaseService) Update(ctx context.Context, model *entity.SysDict) (bool, error) {
	return dao.SysDictDao.UpdateContext(ctx, model,
		meta.SysDictMeta.PrimaryKey(),
		meta.SysDictMeta.Ctime,
		meta.SysDictMeta.Utime,
	)
}

// BatchUpdate 批量更新
func (svc dictBaseService) BatchUpdate(ctx context.Context, param *types.BatchUpdate) (bool, error) {
	for _, row := range param.Rows {
		var id any
		attr := map[string]any{}
		for k, v := range row {
			if k == meta.SysDictMeta.PrimaryKey() {
				id = v
				continue
			}
			attr[k] = v
		}
		err := db.GetDefaultTxManager().ExecTx(ctx, func(txCtx context.Context, tx *sqlx.Tx) error {
			_, err := dao.SysDictDao.UpdateFieldTxContext(txCtx, tx, id, attr)
			return err
		})
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// DeleteByIDs 批量更新
func (svc dictBaseService) DeleteByIDs(ctx context.Context, ids []int64) error {
	l := log.GetLogger(ctx).With(zap.Any("ids", ids))
	_, err := dao.SysDictDao.DeleteByCondContext(ctx, ql.C().And(
		meta.SysDictMeta.IdIn(ids...),
	))
	if err != nil {
		l.Error("delete sys_dict err", zap.Error(err))
		return err
	}
	l.Info("delete sys_dict success")
	return nil
}
