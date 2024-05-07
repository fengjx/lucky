package errno

import "github.com/fengjx/luchen"

// common

var (
	SystemErr       = &luchen.Errno{HTTPCode: 500, Code: 500, Msg: "系统错误"}
	ArgsErr         = &luchen.Errno{HTTPCode: 400, Code: 400, Msg: "参数错误"}
	UnauthorizedErr = &luchen.Errno{HTTPCode: 401, Code: 401, Msg: "Unauthorized"}
	ForbiddenErr    = &luchen.Errno{HTTPCode: 403, Code: 403, Msg: "Forbidden"}
)

// user

var (
	PasswordErr     = &luchen.Errno{HTTPCode: 401, Code: 10000, Msg: "密码错误"}
	UserNotExistErr = &luchen.Errno{HTTPCode: 401, Code: 10001, Msg: "用户不存在"}
)
