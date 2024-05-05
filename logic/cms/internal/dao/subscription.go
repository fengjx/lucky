package dao

import (
	"github.com/fengjx/daox"

	"github.com/fengjx/lucky/logic/cms/internal/data/meta"
)

var CmsSubscriptionDao *cmsSubscriptionDao

func init() {
	CmsSubscriptionDao = newCmsSubscriptionDao()
}

type cmsSubscriptionDao struct {
	*daox.Dao
}

func newCmsSubscriptionDao() *cmsSubscriptionDao {
	inst := &cmsSubscriptionDao{}
	inst.Dao = daox.NewDaoByMeta(meta.CmsSubscriptionMeta)
	return inst
}
