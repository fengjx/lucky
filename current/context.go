package current

import "context"

type (
	adminIDCtxKey struct{}
)

func AdminUID(ctx context.Context) int64 {
	val := ctx.Value(adminIDCtxKey{})
	if val == nil {
		return 0
	}
	if uid, ok := val.(int64); ok {
		return uid
	}
	return 0
}

func WithAdminUID(ctx context.Context, uid int64) context.Context {
	return context.WithValue(ctx, adminIDCtxKey{}, uid)
}
