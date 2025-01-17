package gtclib

func ValueToPointer[T any](v T) *T {
	return &v
}

// вернет nill, если v равна нулевому значению своего типа
// Пример: v string && v == "" --> ValueToPointerZeroValue("") == nil
func ValueToPointerZeroValue[T comparable](v T) *T {
	var zeroValue T
	if v == zeroValue {
		return nil
	}

	return &v
}