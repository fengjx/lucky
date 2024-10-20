package protocol

import (
	"github.com/fengjx/lucky/logic/sys/internal/data/entity"
)

type UserInfo struct {
	UID      int64  `json:"uid,omitempty"`
	Username string `json:"username,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Email    string `json:"email,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Phone    string `json:"phone,omitempty"`
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
