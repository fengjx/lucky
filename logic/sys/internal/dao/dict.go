package dao

import (
	"context"

	"github.com/fengjx/daox"
	"github.com/fengjx/daox/sqlbuilder/ql"

	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/internal/data/enum"
	"github.com/fengjx/lucky/logic/sys/internal/data/meta"
)

var SysDictDao *sysDictDao

func init() {
	SysDictDao = newSysDictDao()
}

type sysDictDao struct {
	*daox.Dao
}

func newSysDictDao() *sysDictDao {
	inst := &sysDictDao{}
	inst.Dao = daox.NewDaoByMeta(meta.SysDictMeta)
	return inst
}

// ListAll 查询所有生效数据字典
func (dao sysDictDao) ListAll(ctx context.Context) ([]*entity.SysDict, error) {
	var list []*entity.SysDict
	err := dao.Selector().
		Where(ql.C(ql.Col(meta.SysUserMeta.Status).EQ(enum.DictStatusNormal))).
		SelectContext(ctx, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
