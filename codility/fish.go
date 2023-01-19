package codility

import "fmt"

func GreetFish() {
	A := []int{4, 3, 2, 1, 5}
	B := []int{0, 1, 0, 0, 0}
	n := greetFish(A, B)
	fmt.Println(n)

}

func greetFish(A, B []int) int {
	if len(A) == 0 {
		return 0
	}

	live := len(A)
	stack := make([]int, 0)
	for i := 0; i < len(A); i++ {
		if B[i] == 1 {
			stack = append(stack, A[i])
		}
		if B[i] == 0 {
			for !(len(stack) == 0) {
				if stack[len(stack)-1] > A[i] {
					live--
					break
				} else if stack[len(stack)-1] < A[i] {
					live--
					stack = append([]int{}, stack[:len(stack)-1]...)
				}
			}
		}
	}
	return live
}
