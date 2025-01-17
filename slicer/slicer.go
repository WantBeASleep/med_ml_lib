package slicer

func Flatten2DArray[T any](slice [][]T) []T {
	cnt := 0
	for _, v := range slice {
		cnt += len(v)
	}

	res := make([]T, 0, cnt)
	for _, v := range slice {
		res = append(res, v...)
	}

	return res
}


func PackSlice[T any](slice []T) []any {
	res := make([]any, 0, len(slice))
	for _, v := range slice {
		res = append(res, any(v))
	}

	return res
}