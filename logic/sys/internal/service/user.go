package service

import (
	"context"
	"strings"

	"github.com/fengjx/go-halo/errs"
	"github.com/fengjx/go-halo/utils"

	"github.com/fengjx/lucky/logic/sys/internal/dao"
	"github.com/fengjx/lucky/logic/sys/internal/dao/schema"
	"github.com/fengjx/lucky/pkg/kit"
)

var UserSvc = &userService{}

type userService struct {
}

func (s *userService) GetByUsername(ctx context.Context, username string) (*schema.SysUser, error) {
	user, err := dao.SysUserDao.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdatePwd 修改用户密码
func (s *userService) UpdatePwd(ctx context.Context, id int64, newPwd string) error {
	pwd, salt := s.genPwd(newPwd)
	_, err := dao.SysUserDao.UpdatePwd(ctx, id, pwd, salt)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) genPwd(pwd string) (md5Pwd, salt string) {
	if pwd == "" {
		return
	}
	salt = utils.RandomString(6)
	sb := strings.Builder{}
	sb.WriteString(pwd)
	sb.WriteString(salt)
	md5Pwd = kit.MD5Hash(sb.String())
	return
}

func (s *userService) Get(ctx context.Context, uid int64) (*schema.SysUser, error) {
	if uid == 0 {
		return nil, nil
	}
	user, err := dao.SysUserDao.GetByIDContext(ctx, uid)
	if err != nil {
		return nil, errs.Wrap(err, "get user err")
	}
	return user, nil
}
