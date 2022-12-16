package utils

func Contains() {
	arr := []int{1, 2, 3, 4, 5}
	v := 3

	res := contains[int](arr, v)
}

func contains[T string | int](arr []T, v T) bool {
	for i := range arr {
		if arr[i] == v {
			return true
		}
	}
	return false
}
