package endpoint

import (
	"context"
	"strings"

	"github.com/fengjx/luchen"
	"github.com/fengjx/luchen/log"
	"go.uber.org/zap"

	"github.com/fengjx/lucky/connom/auth"
	"github.com/fengjx/lucky/connom/errno"
	"github.com/fengjx/lucky/connom/kit"
	"github.com/fengjx/lucky/current"
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
	"github.com/fengjx/lucky/logic/sys/internal/protocol"
	"github.com/fengjx/lucky/logic/sys/internal/service"
)

var login = loginEndpoint{}

type loginEndpoint struct {
}

func (e loginEndpoint) makeLoginEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*protocol.LoginReq)
		l := log.GetLogger(ctx).With(zap.String("username", req.Username))
		user, err := service.UserSvc.GetByUsername(ctx, req.Username)
		if err != nil {
			l.Error("user login err",
				zap.String("username", req.Username),
				zap.Error(err),
			)
			return nil, err
		}
		if user == nil {
			l.Warn("user not exist", zap.String("username", req.Username))
			return nil, errno.UserNotExistErr
		}
		if !checkPassword(user, req.Password) {
			return nil, errno.PasswordErr
		}
		token, err := auth.GenToken(auth.LoginPayload{
			UID: user.ID,
		})
		if err != nil {
			l.Error("gen token err", zap.Error(err))
			return nil, err
		}
		resp := &protocol.LoginResp{
			Token: token,
		}
		resp.UserInfo = protocol.BuildUserInfo(user)
		return resp, nil
	}
}

func (e loginEndpoint) makeUserInfoEndpoint() luchen.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		uid := current.UID(ctx)
		user, err := service.UserBaseSvc.Get(ctx, uid)
		if err != nil {
			log.ErrorCtx(ctx, "get user_info err", zap.Int64("uid", uid), zap.Error(err))
			return nil, err
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
