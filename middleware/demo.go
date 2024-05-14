package middleware

import (
	"context"

	"github.com/fengjx/luchen"
	"github.com/fengjx/luchen/env"

	"github.com/fengjx/lucky/connom/errno"
)

var DemoForbiddenMiddleware = ForbiddenMiddleware("local")

// ForbiddenMiddleware 指定环境禁止执行
func ForbiddenMiddleware(envName string) luchen.Middleware {
	return func(next luchen.Endpoint) luchen.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if env.Is(envName) {
				return nil, errno.ForbiddenErr
			}
			return
		}
	}
}
