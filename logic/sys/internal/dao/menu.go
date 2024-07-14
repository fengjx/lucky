package dao

import (
	"context"

	"github.com/fengjx/daox"
	"github.com/fengjx/daox/sqlbuilder/ql"
	"github.com/samber/lo"

	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/internal/data/enum"
	"github.com/fengjx/lucky/logic/sys/internal/data/meta"
)

var SysMenuDao *sysMenuDao

func init() {
	SysMenuDao = newSysMenuDao()
}

type sysMenuDao struct {
	*daox.Dao
}

func newSysMenuDao() *sysMenuDao {
	inst := &sysMenuDao{}
	inst.Dao = daox.NewDaoByMeta(meta.SysMenuMeta)
	return inst
}

// ListChildren 查询所有子菜单
func (dao sysMenuDao) ListChildren(ctx context.Context, status []enum.MenuStatus, parentIDs ...int64) (map[int64][]*entity.SysMenu, error) {
	var list []*entity.SysMenu
	err := dao.Selector().
		Where(
			ql.C(
				meta.SysMenuMeta.StatusIn(lo.Map[enum.MenuStatus, string](status, func(item enum.MenuStatus, index int) string {
					return string(item)
				})...),
				meta.SysMenuMeta.ParentIdIn(parentIDs...),
			),
		).OrderBy(meta.SysMenuMeta.SortNoAsc()).
		SelectContext(ctx, &list)
	if err != nil {
		return nil, err
	}
	resMap := map[int64][]*entity.SysMenu{}
	for _, menu := range list {
		resMap[menu.ParentID] = append(resMap[menu.ParentID], menu)
	}
	return resMap, nil
}
