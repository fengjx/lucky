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

var UserBaseSvc = &userBaseService{}

type userBaseService struct {
}

// Query 分页查询
func (svc userBaseService) Query(ctx context.Context, query *daox.QueryRecord) (*types.PageVO[entity.CmsUser], error) {
	readDB := dao.CmsUserDao.GetReadDB()
	query.TableName = meta.CmsUserMeta.TableName()
	list, page, err := daox.Find[entity.CmsUser](ctx, readDB, *query)
	if err != nil {
		log.ErrorCtx(ctx, "page query cms_user err", zap.Any("query", json.ToJsonDelay(query)), zap.Error(err))
		return nil, err
	}
	pageVO := &types.PageVO[entity.CmsUser]{
		List:    list,
		Offset:  page.Offset,
		Limit:   page.Limit,
		Count:   page.Count,
		HasNext: page.HasNext,
	}
	return pageVO, nil
}

// Add 新增记录
func (svc userBaseService) Add(ctx context.Context, model *entity.CmsUser) (int64, error) {
	return dao.CmsUserDao.SaveContext(ctx, model)
}

// Update 更新记录
func (svc userBaseService) Update(ctx context.Context, model *entity.CmsUser) (bool, error) {
	return dao.CmsUserDao.UpdateContext(ctx, model,
		meta.CmsUserMeta.PrimaryKey(),
	)
}

// BatchUpdate 批量更新
func (svc userBaseService) BatchUpdate(ctx context.Context, param *types.BatchUpdate) (bool, error) {
	for _, row := range param.Rows {
		var id any
		attr := map[string]any{}
		for k, v := range row {
			if k == meta.CmsUserMeta.PrimaryKey() {
				id = v
				continue
			}
			attr[k] = v
		}
		err := db.GetDefaultTxManager().ExecTx(ctx, func(txCtx context.Context, tx *sqlx.Tx) error {
			_, err := dao.CmsUserDao.UpdateFieldTxContext(txCtx, tx, id, attr)
			return err
		})
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// DeleteByIDs 批量更新
func (svc userBaseService) DeleteByIDs(ctx context.Context, ids []int64) error {
	l := log.GetLogger(ctx).With(zap.Any("ids", ids))
	_, err := dao.CmsUserDao.Deleter().Where(
		ql.C(
			meta.CmsUserMeta.IdIn(ids...),
		),
	).ExecContext(ctx)
	if err != nil {
		l.Error("delete cms_user err", zap.Error(err))
		return err
	}
	l.Info("delete cms_user success")
	return nil
}
