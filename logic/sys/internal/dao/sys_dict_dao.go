package dao

import (
	"context"

	"github.com/fengjx/daox/v2"

	"github.com/fengjx/lucky/integration/db"
	"github.com/fengjx/lucky/logic/sys/internal/dao/schema"
	"github.com/fengjx/lucky/logic/sys/internal/dao/schema/sysdict"
	"github.com/fengjx/lucky/logic/sys/internal/data/enum"
)

var SysDictDao *sysDictDao

func init() {
	SysDictDao = newSysDictDao()
}

type sysDictDao struct {
	*daox.Dao[*schema.SysDict]
}

func newSysDictDao() *sysDictDao {
	dao := daox.NewDao[*schema.SysDict](
		sysdict.Meta,
		daox.WithDBMaster(db.GetDefaultDB()),
	)
	inst := &sysDictDao{
		Dao: dao,
	}
	return inst
}

// ListAll 查询所有生效数据字典
func (d *sysDictDao) ListAll(ctx context.Context) ([]*schema.SysDict, error) {
	var list []*schema.SysDict
	err := d.Selector().
		WhereC(
			sysdict.StatusEQ(string(enum.DictStatusNormal)),
		).
		ListContext(ctx, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
