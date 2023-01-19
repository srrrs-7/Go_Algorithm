package codility

import "fmt"

func FlogJump() {
	X, Y, D := 5, 105, 3
	n := flogJump(X, Y, D)
	fmt.Println(n)
}

func flogJump(X int, Y int, D int) int {
	distance := Y - X
	reach := distance / D
	if distance%D == 0 {
		return reach
	}
	return reach + 1
}
