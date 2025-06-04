package dao

import (
	"context"

	"github.com/fengjx/daox/v2"
	"github.com/fengjx/go-halo/errs"

	"github.com/fengjx/lucky/integration/db"
	"github.com/fengjx/lucky/logic/sys/internal/dao/schema"
	"github.com/fengjx/lucky/logic/sys/internal/dao/schema/sysuser"
)

var SysUserDao *sysUserDao

func init() {
	SysUserDao = newSysUserDao()
}

type sysUserDao struct {
	*daox.Dao[*schema.SysUser]
}

func newSysUserDao() *sysUserDao {
	dao := daox.NewDao[*schema.SysUser](
		sysuser.Meta,
		daox.WithDBMaster(db.GetDefaultDB()),
	)
	inst := &sysUserDao{
		Dao: dao,
	}
	return inst
}

// GetByUsername 根据用户名获取用户
func (d *sysUserDao) GetByUsername(ctx context.Context, username string) (*schema.SysUser, error) {
	user, err := d.GetByColumnContext(ctx, daox.OfKv(sysuser.UsernameField, username))
	if err != nil {
		return nil, errs.Wrap(err, "get user by username err")
	}
	return user, nil
}

// UpdatePwd 修改用户密码
func (d *sysUserDao) UpdatePwd(ctx context.Context, id int64, pwd string, salt string) (bool, error) {
	affected, err := d.UpdateFieldContext(ctx, id, map[string]any{
		sysuser.PwdField:  pwd,
		sysuser.SaltField: salt,
	})
	if err != nil {
		return false, errs.Wrap(err, "update user password err")
	}
	return affected > 0, nil
}
