package leetcode

import (
	"fmt"
	"math"
)

func MinimumCostAve() {
	nums := []int{1, 3, 5, 2}
	costs := []int{2, 3, 1, 14}

	res := minimumCostAve(nums, costs)
	fmt.Println("minimum cost: ", res)
}

func minimumCostAve(nums []int, costs []int) int {
	sum := 0
	for i := range nums {
		sum += i
	}
	ave := sum / len(nums)

	minCost := 0
	for i, n := range nums {
		minCost += int(math.Abs(float64(n-ave))) * costs[i]
		fmt.Println(ave)
	}

	return minCost
}
