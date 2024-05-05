package protocol

import (
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
)

type UserInfo struct {
	UID      int64  `json:"uid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
}

func BuildUserInfo(e *entity.SysUser) *UserInfo {
	if e == nil {
		return nil
	}
	return &UserInfo{
		UID:      e.ID,
		Username: e.Username,
		Nickname: e.Nickname,
		Email:    e.Email,
		Avatar:   e.Avatar,
		Phone:    e.Phone,
	}
}

type UpdateUserPwdReq struct {
	ID  int64  `json:"id"`
	Pwd string `json:"pwd"`
}
