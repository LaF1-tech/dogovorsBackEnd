package uslices

func MapFunc[T any, R any](arr []T, converter func(item T) R) []R {
	newArr := make([]R, len(arr))
	for i, item := range arr {
		newArr[i] = converter(item)
	}
	return newArr
}
