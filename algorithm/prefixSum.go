package algorithm

import "fmt"

func PrefixSum() {
	arr := []int{10, 20, 30, 40, 50}
	arr2 := prefixSum(arr)
	fmt.Println(arr2)

	n := rangeSum(arr, 0, 4)
	fmt.Println("range sum", n)
}

func prefixSum(arr []int) []int {
	var arr2 = make([]int, len(arr)-1)
	arr2[0] = arr[0]
	for i := 1; i < len(arr2); i++ {
		arr2[i] += arr2[i-1] + arr[i]
	}
	return arr2
}

func rangeSum(arr []int, from, to int) int {
	var num int
	for i := from; i <= to; i++ {
		num += arr[i]
	}
	return num
}
