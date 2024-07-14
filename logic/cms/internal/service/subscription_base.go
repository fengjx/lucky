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

var SubscriptionBaseSvc = &subscriptionBaseService{}

type subscriptionBaseService struct {
}

// Query 分页查询
func (svc subscriptionBaseService) Query(ctx context.Context, query *daox.QueryRecord) (*types.PageVO[entity.CmsSubscription], error) {
	readDB := dao.CmsSubscriptionDao.GetReadDB()
	query.TableName = meta.CmsSubscriptionMeta.TableName()
	list, page, err := daox.Find[entity.CmsSubscription](ctx, readDB, *query)
	if err != nil {
		log.ErrorCtx(ctx, "page query cms_subscription err", zap.Any("query", json.ToJsonDelay(query)), zap.Error(err))
		return nil, err
	}
	pageVO := &types.PageVO[entity.CmsSubscription]{
		List:    list,
		Offset:  page.Offset,
		Limit:   page.Limit,
		Count:   page.Count,
		HasNext: page.HasNext,
	}
	return pageVO, nil
}

// Add 新增记录
func (svc subscriptionBaseService) Add(ctx context.Context, model *entity.CmsSubscription) (int64, error) {
	return dao.CmsSubscriptionDao.SaveContext(ctx, model)
}

// Update 更新记录
func (svc subscriptionBaseService) Update(ctx context.Context, model *entity.CmsSubscription) (bool, error) {
	return dao.CmsSubscriptionDao.UpdateContext(ctx, model,
		meta.CmsSubscriptionMeta.PrimaryKey(),
	)
}

// BatchUpdate 批量更新
func (svc subscriptionBaseService) BatchUpdate(ctx context.Context, param *types.BatchUpdate) (bool, error) {
	for _, row := range param.Rows {
		var id any
		attr := map[string]any{}
		for k, v := range row {
			if k == meta.CmsSubscriptionMeta.PrimaryKey() {
				id = v
				continue
			}
			attr[k] = v
		}
		err := db.GetDefaultTxManager().ExecTx(ctx, func(txCtx context.Context, tx *sqlx.Tx) error {
			_, err := dao.CmsSubscriptionDao.UpdateFieldTxContext(txCtx, tx, id, attr)
			return err
		})
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// DeleteByIDs 批量更新
func (svc subscriptionBaseService) DeleteByIDs(ctx context.Context, ids []int64) error {
	l := log.GetLogger(ctx).With(zap.Any("ids", ids))
	_, err := dao.CmsSubscriptionDao.Deleter().Where(
		ql.C(
			meta.CmsSubscriptionMeta.IdIn(ids...),
		),
	).ExecContext(ctx)
	if err != nil {
		l.Error("delete cms_subscription err", zap.Error(err))
		return err
	}
	l.Info("delete cms_subscription success")
	return nil
}
