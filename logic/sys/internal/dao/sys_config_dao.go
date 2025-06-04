package dao

import (
	"context"

	"github.com/fengjx/daox/v2"
	"github.com/fengjx/go-halo/errs"

	"github.com/fengjx/lucky/integration/db"
	"github.com/fengjx/lucky/logic/sys/internal/dao/schema"
	"github.com/fengjx/lucky/logic/sys/internal/dao/schema/sysconfig"
	"github.com/fengjx/lucky/logic/sys/internal/data/enum"
	"github.com/fengjx/lucky/pkg/kit"
)

var SysConfigDao *sysConfigDao

func init() {
	SysConfigDao = newSysConfigDao()
}

type sysConfigDao struct {
	*daox.Dao[*schema.SysConfig]
}

func newSysConfigDao() *sysConfigDao {
	dao := daox.NewDao[*schema.SysConfig](
		sysconfig.Meta,
		daox.WithDBMaster(db.GetDefaultDB()),
	)
	inst := &sysConfigDao{
		Dao: dao,
	}
	return inst
}

// ListAll 查询所有生效配置
func (d *sysConfigDao) ListAll(ctx context.Context) ([]*schema.SysConfig, error) {
	var list []*schema.SysConfig
	err := d.Selector().
		WhereC(
			sysconfig.StatusEQ(string(enum.ConfigStatusNormal)),
		).
		ListContext(ctx, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// ListScopeConfig 查询指定作用域的配置
func (d *sysConfigDao) ListScopeConfig(ctx context.Context, scopes []string) ([]*schema.SysConfig, error) {
	list, err := d.ListByColumnsContext(ctx, daox.OfMultiKv(sysconfig.ScopeField, kit.ToAnySlice(scopes)...))
	if err != nil {
		return nil, errs.Wrap(err, "list all sys_config err")
	}
	return list, nil
}
