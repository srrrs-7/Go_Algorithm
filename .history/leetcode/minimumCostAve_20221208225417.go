package leetcode

import (
	"fmt"
	"math"
)

func MinimumCostAve() {
	nums := []int{2, 2, 2, 2}
	costs := []int{2, 3, 1, 14}

	res := minimumCostAve(nums, costs)
	fmt.Println("minimum cost: ", res)
}

func minimumCostAve(nums []int, costs []int) int {
	sum := 0
	for i := range nums {
		sum += i
	}
	ave := sum / (len(nums) - 1)

	minCost := 0
	for i, n := range nums {
		minCost += int(math.Abs(float64(n-ave))) * costs[i]
	}

	return minCost
}
