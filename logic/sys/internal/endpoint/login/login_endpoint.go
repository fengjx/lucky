package login

import (
	"context"
	"reflect"
	"strings"

	"github.com/fengjx/go-halo/errs"
	"github.com/fengjx/luchen"
	"github.com/fengjx/lucky/common/auth"
	"github.com/fengjx/lucky/common/errno"
	"github.com/fengjx/lucky/common/kit"
	"github.com/fengjx/lucky/common/types"
	"github.com/fengjx/lucky/current"
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/internal/protocol"
	"github.com/fengjx/lucky/logic/sys/internal/service"
)

func RegisterLoginTTPHandler(hs *luchen.HTTPServer) {
	e := &loginEndpoint{}
	hs.Handle(&luchen.EndpointDefine{
		Name:     "AdminLogin.Login",
		Path:     "/api/open/sys/login",
		ReqType:  reflect.TypeOf(&protocol.LoginReq{}),
		RspType:  reflect.TypeOf(&protocol.LoginResp{}),
		Endpoint: e.makeLoginEndpoint(),
	})

	hs.Handle(&luchen.EndpointDefine{
		Name:     "AdminLogin.UserInfo",
		Path:     "/admin/sys/user/info",
		ReqType:  reflect.TypeOf(&types.Empty{}),
		RspType:  reflect.TypeOf(&protocol.UserInfoResp{}),
		Endpoint: e.makeUserInfoEndpoint(),
	})
}

type loginEndpoint struct {
}

func (e *loginEndpoint) makeLoginEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*protocol.LoginReq)
		user, err := service.UserSvc.GetByUsername(ctx, req.Username)
		if err != nil {
			return nil, errs.Wrap(err, "user login err")
		}
		if user == nil {
			return nil, errs.WithStack(errno.ErrUserNotExist)
		}
		if !checkPassword(user, req.Password) {
			return nil, errno.ErrPassword
		}
		token, err := auth.GenToken(auth.LoginPayload{
			UID: user.ID,
		})
		if err != nil {
			return nil, errs.Wrap(err, "gen token err")
		}
		resp := &protocol.LoginResp{
			Token: token,
		}
		resp.UserInfo = protocol.BuildUserInfo(user)
		return resp, nil
	}
}

func (e *loginEndpoint) makeUserInfoEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		uid := current.AdminUID(ctx)
		user, err := service.UserBaseSvc.Get(ctx, uid)
		if err != nil {
			return nil, errs.Wrap(err, "get user_info err")
		}
		resp := &protocol.UserInfoResp{
			UserInfo: protocol.BuildUserInfo(user),
		}
		return resp, nil
	}
}

// checkPassword 检查密码是否匹配
func checkPassword(user *entity.SysUser, password string) bool {
	sb := strings.Builder{}
	sb.WriteString(password)
	sb.WriteString(user.Salt)
	md5Pwd := kit.MD5Hash(sb.String())
	return user.Pwd == md5Pwd
}
