package algods

// Swap values with provided indexes in array.
func Swap[T any](arr []T, ia, ib int) {
	if arr == nil || ia > len(arr) || ib > len(arr) {
		return
	}

	arr[ia], arr[ib] = arr[ib], arr[ia]
}
