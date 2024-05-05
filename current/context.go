package current

import "context"

type (
	uidCtxKey struct{}
)

func UID(ctx context.Context) int64 {
	val := ctx.Value(uidCtxKey{})
	if val == nil {
		return 0
	}
	if uid, ok := val.(int64); ok {
		return uid
	}
	return 0
}

func WithUID(ctx context.Context, uid int64) context.Context {
	return context.WithValue(ctx, uidCtxKey{}, uid)
}
