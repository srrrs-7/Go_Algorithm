package atCorder

import "fmt"

func Dp() {
	arr := []int{1, 4, 9}
	t := 7

	res := dp(arr, t)
	fmt.Println("multi number count: ", res)
}

// digit DP
func dp(arr []int, t int) int {

}