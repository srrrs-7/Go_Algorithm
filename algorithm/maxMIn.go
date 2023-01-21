package algorithm

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mix(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Solution(A []int) int {
	min := A[0]
	max := A[0]
	count := 1
	maxCount := 1
	for i := 1; i < len(A); i++ {
		if A[i] <= max && A[i] >= min {
			count++
		} else if A[i] > max {
			max = A[i]
			count = 2
		} else if A[i] < min {
			min = A[i]
			count = 2
		}
		maxCount = maxInt(maxCount, count)
	}
	return maxCount
}
