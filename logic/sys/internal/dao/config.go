package dao

import (
	"context"

	"github.com/fengjx/daox"
	"github.com/fengjx/daox/sqlbuilder/ql"

	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/internal/data/enum"
	"github.com/fengjx/lucky/logic/sys/internal/data/meta"
)

var SysConfigDao *sysConfigDao

func init() {
	SysConfigDao = newSysConfigDao()
}

type sysConfigDao struct {
	*daox.Dao
}

func newSysConfigDao() *sysConfigDao {
	inst := &sysConfigDao{}
	inst.Dao = daox.NewDaoByMeta(meta.SysConfigMeta)
	return inst
}

// ListAll 查询所有生效配置
func (dao sysConfigDao) ListAll(ctx context.Context) ([]*entity.SysConfig, error) {
	selector := dao.Selector().Where(ql.C().And(
		ql.Col(meta.SysUserMeta.Status).EQ(enum.ConfigStatusNormal),
	))
	var list []*entity.SysConfig
	err := dao.SelectContext(ctx, &list, selector)
	if err != nil {
		return nil, err
	}
	return list, nil
}
