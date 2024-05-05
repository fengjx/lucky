package dao

import (
	"github.com/fengjx/daox"

	"github.com/fengjx/lucky/logic/sys/internal/data/meta"
)

var SysUserDao *sysUserDao

func init() {
	SysUserDao = newSysUserDao()
}

type sysUserDao struct {
	*daox.Dao
}

func newSysUserDao() *sysUserDao {
	inst := &sysUserDao{}
	inst.Dao = daox.NewDaoByMeta(meta.SysUserMeta)
	return inst
}
