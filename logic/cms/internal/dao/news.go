package dao

import (
	"github.com/fengjx/daox"

	"github.com/fengjx/lucky/logic/cms/internal/data/meta"
)

var CmsNewsDao *cmsNewsDao

func init() {
	CmsNewsDao = newCmsNewsDao()
}

type cmsNewsDao struct {
	*daox.Dao
}

func newCmsNewsDao() *cmsNewsDao {
	inst := &cmsNewsDao{}
	inst.Dao = daox.NewDaoByMeta(meta.CmsNewsMeta)
	return inst
}
