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

var ConfigBaseSvc = &configBaseService{}

type configBaseService struct {
}

// Query 分页查询
func (svc configBaseService) Query(ctx context.Context, query *daox.QueryRecord) (*types.PageVO[entity.SysConfig], error) {
	readDB := dao.SysConfigDao.GetReadDB()
	query.TableName = meta.SysConfigMeta.TableName()
	list, page, err := daox.Find[entity.SysConfig](ctx, readDB, *query)
	if err != nil {
		log.ErrorCtx(ctx, "page query sys_config err", zap.Any("query", json.ToJsonDelay(query)), zap.Error(err))
		return nil, err
	}
	pageVO := &types.PageVO[entity.SysConfig]{
		List:    list,
		Offset:  page.Offset,
		Limit:   page.Limit,
		Count:   page.Count,
		HasNext: page.HasNext,
	}
	return pageVO, nil
}

// Add 新增记录
func (svc configBaseService) Add(ctx context.Context, model *entity.SysConfig) (int64, error) {
	return dao.SysConfigDao.SaveContext(ctx, model,
		daox.WithInsertOmits(
			meta.SysConfigMeta.Ctime,
			meta.SysConfigMeta.Utime,
		),
	)
}

// Update 更新记录
func (svc configBaseService) Update(ctx context.Context, model *entity.SysConfig) (bool, error) {
	return dao.SysConfigDao.UpdateContext(ctx, model,
		meta.SysConfigMeta.PrimaryKey(),
		meta.SysConfigMeta.Ctime,
		meta.SysConfigMeta.Utime,
	)
}

// BatchUpdate 批量更新
func (svc configBaseService) BatchUpdate(ctx context.Context, param *types.BatchUpdate) (bool, error) {
	for _, row := range param.Rows {
		var id any
		attr := map[string]any{}
		for k, v := range row {
			if k == meta.SysConfigMeta.PrimaryKey() {
				id = v
				continue
			}
			attr[k] = v
		}
		err := db.GetDefaultTxManager().ExecTx(ctx, func(txCtx context.Context, tx *sqlx.Tx) error {
			_, err := dao.SysConfigDao.UpdateFieldTxContext(txCtx, tx, id, attr)
			return err
		})
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// DeleteByIDs 批量更新
func (svc configBaseService) DeleteByIDs(ctx context.Context, ids []int64) error {
	l := log.GetLogger(ctx).With(zap.Any("ids", ids))
	_, err := dao.SysConfigDao.Deleter().
		Where(ql.C(meta.SysConfigMeta.IdIn(ids...))).
		ExecContext(ctx)
	if err != nil {
		l.Error("delete sys_config err", zap.Error(err))
		return err
	}
	l.Info("delete sys_config success")
	return nil
}
