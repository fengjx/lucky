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
	"github.com/fengjx/lucky/logic/cms/internal/dao"
	"github.com/fengjx/lucky/logic/cms/internal/data/entity"
	"github.com/fengjx/lucky/logic/cms/internal/data/meta"
)

var NewsBaseSvc = &newsBaseService{}

type newsBaseService struct {
}

// Query 分页查询
func (svc newsBaseService) Query(ctx context.Context, query *daox.QueryRecord) (*types.PageVO[entity.CmsNews], error) {
	readDB := dao.CmsNewsDao.GetReadDB()
	query.TableName = meta.CmsNewsMeta.TableName()
	list, page, err := daox.Find[entity.CmsNews](ctx, readDB, *query)
	if err != nil {
		log.ErrorCtx(ctx, "page query cms_news err", zap.Any("query", json.ToJsonDelay(query)), zap.Error(err))
		return nil, err
	}
	pageVO := &types.PageVO[entity.CmsNews]{
		List:    list,
		Offset:  page.Offset,
		Limit:   page.Limit,
		Count:   page.Count,
		HasNext: page.HasNext,
	}
	return pageVO, nil
}

// Add 新增记录
func (svc newsBaseService) Add(ctx context.Context, model *entity.CmsNews) (int64, error) {
	return dao.CmsNewsDao.SaveContext(ctx, model)
}

// Update 更新记录
func (svc newsBaseService) Update(ctx context.Context, model *entity.CmsNews) (bool, error) {
	return dao.CmsNewsDao.UpdateContext(ctx, model,
		meta.CmsNewsMeta.PrimaryKey(),
		meta.CmsNewsMeta.Utime,
		meta.CmsNewsMeta.Ctime,
	)
}

// BatchUpdate 批量更新
func (svc newsBaseService) BatchUpdate(ctx context.Context, param *types.BatchUpdate) (bool, error) {
	for _, row := range param.Rows {
		var id any
		attr := map[string]any{}
		for k, v := range row {
			if k == meta.CmsNewsMeta.PrimaryKey() {
				id = v
				continue
			}
			attr[k] = v
		}
		err := db.GetDefaultTxManager().ExecTx(ctx, func(txCtx context.Context, tx *sqlx.Tx) error {
			_, err := dao.CmsNewsDao.UpdateFieldTxContext(txCtx, tx, id, attr)
			return err
		})
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// DeleteByIDs 批量更新
func (svc newsBaseService) DeleteByIDs(ctx context.Context, ids []int64) error {
	l := log.GetLogger(ctx).With(zap.Any("ids", ids))
	_, err := dao.CmsNewsDao.DeleteByCondContext(ctx, ql.C().And(
		meta.CmsNewsMeta.IdIn(ids...),
	))
	if err != nil {
		l.Error("delete cms_news err", zap.Error(err))
		return err
	}
	l.Info("delete cms_news success")
	return nil
}
