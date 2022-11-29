package atCorder

import "fmt"

func Dp() {
	arr := []int{1, 4, 9}
	multi := 7

	res := dp(arr, multi)
	fmt.Println("multi number count: ", res)
}

// digit DP
func dp(arr []int, multi int) int {

}