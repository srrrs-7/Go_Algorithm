package algorithm

import "fmt"

func MaxSubArray() {
	arr := []int{5, 4, -1, 7, 8}
	n := maxSubArray(arr)
	fmt.Println("max sub array", n)
}

func maxSubArray(arr []int) int {
	curSum := arr[0]
	maxSum := arr[0]

	for i := 1; i < len(arr); i++ {
		curSum = maxInt(curSum+arr[i], arr[i])
		maxSum = maxInt(curSum, maxSum)
	}

	return maxSum
}
