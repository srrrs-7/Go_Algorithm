package codility

import (
	"fmt"
	"math"
)

func MaxProfit() {
	arr := []int{23171, 21011, 21123, 21366, 21013, 21367}
	n := maxProfit(arr)
	fmt.Println(n)
}

func maxProfit(A []int) int {
	var maxProfit = 0
	var minPrice = math.MaxUint32

	for i := 0; i < len(A); i++ {
		minPrice = minInt(minPrice, A[i])
		maxProfit = maxInt(maxProfit, A[i]-minPrice)

		fmt.Println(minPrice, maxProfit)
	}
	return maxProfit
}

func minInt(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
