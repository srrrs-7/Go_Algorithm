package utils

func Contains() {}

func contains[T string | int](arr []T, v T) bool {
	for i := range arr {
		if arr[i] == v {
			return true
		}
	}
	return false
}
