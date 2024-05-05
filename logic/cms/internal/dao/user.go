package dao

import (
	"github.com/fengjx/daox"

	"github.com/fengjx/lucky/logic/cms/internal/data/meta"
)

var CmsUserDao *cmsUserDao

func init() {
	CmsUserDao = newCmsUserDao()
}

type cmsUserDao struct {
	*daox.Dao
}

func newCmsUserDao() *cmsUserDao {
	inst := &cmsUserDao{}
	inst.Dao = daox.NewDaoByMeta(meta.CmsUserMeta)
	return inst
}
