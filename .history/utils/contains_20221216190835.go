package utils

func Contains() {
	arr := []int{1, 2, 3, 45}
}

func contains[T string | int](arr []T, v T) bool {
	for i := range arr {
		if arr[i] == v {
			return true
		}
	}
	return false
}