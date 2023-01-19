package algorithm

func pop(arr *[]int) (int, []int) {
	x := (*arr)[len(*arr)-1]    // Store the last value to return.
	*arr = (*arr)[:len(*arr)-1] // Remove the last value from the slice.
	return x, *arr
}
