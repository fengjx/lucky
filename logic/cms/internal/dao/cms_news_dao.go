package dao

import (
	"github.com/fengjx/daox/v2"

	"github.com/fengjx/lucky/integration/db"
	"github.com/fengjx/lucky/logic/cms/internal/dao/schema"
	"github.com/fengjx/lucky/logic/cms/internal/dao/schema/cmsnews"
)

var CmsNewsDao *cmsNewsDao

func init() {
	CmsNewsDao = newCmsNewsDao()
}

type cmsNewsDao struct {
	*daox.Dao[*schema.CmsNews]
}

func newCmsNewsDao() *cmsNewsDao {
	dao := daox.NewDao[*schema.CmsNews](
		cmsnews.Meta,
		daox.WithDBMaster(db.GetDefaultDB()),
	)
	inst := &cmsNewsDao{
		Dao: dao,
	}
	return inst
}
