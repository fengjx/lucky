package kit

func ToAnySlice[T any](src []T) []any {
	var arr []any
	for _, item := range src {
		arr = append(arr, item)
	}
	return arr
}
