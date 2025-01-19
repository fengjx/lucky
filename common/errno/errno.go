package errno

import (
	"net/http"

	"github.com/fengjx/luchen"
)

// 用户相关错误码
var (
	ErrPassword     = luchen.NewErrno(http.StatusBadRequest, 1000, "密码错误")
	ErrUserNotExist = luchen.NewErrno(http.StatusNotFound, 1001, "用户不存在")
)
