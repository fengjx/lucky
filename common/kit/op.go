package kit

// OrElse 代替三目运算
func OrElse[T any](cond bool, value T, other T) T {
	if cond {
		return value
	}
	return other
}
