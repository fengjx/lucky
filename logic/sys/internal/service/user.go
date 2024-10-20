package service

import (
	"context"

	"github.com/fengjx/daox"
	"github.com/fengjx/go-halo/errs"

	"github.com/fengjx/lucky/logic/sys/internal/dao"
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/internal/data/meta"
)

var UserSvc = &userService{}

type userService struct {
}

func (svc userService) GetByUsername(ctx context.Context, username string) (*entity.SysUser, error) {
	user := &entity.SysUser{}
	ok, err := dao.SysUserDao.GetByColumnContext(ctx, daox.OfKv(meta.SysUserMeta.Username, username), user)
	if err != nil {
		return nil, errs.Wrap(err, "get user by username err")
	}
	if !ok {
		return nil, nil
	}
	return user, nil
}
