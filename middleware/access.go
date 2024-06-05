package middleware

import (
	"context"

	"github.com/fengjx/luchen"

	"github.com/fengjx/lucky/current"
)

// AccessMiddleware 请求日志
var AccessMiddleware = luchen.AccessMiddleware(&luchen.AccessLogOpt{
	MaxDay: 15,
	ContextFields: map[string]luchen.GetValueFromContext{
		"uid": func(ctx context.Context) any {
			return current.UID(ctx)
		},
	},
})
