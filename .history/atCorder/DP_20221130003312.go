package atCorder

import "fmt"

func Dp() {
	l := []int{1, 4, 9}
	t := 7

	res := dp(l, t)
	fmt.Println("multi number count: ", res)
}

// digit DP
func dp(l []int, t int) int {

}