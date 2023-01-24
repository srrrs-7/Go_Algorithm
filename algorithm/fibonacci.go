package algorithm

import "fmt"

func Fibonacci() {
	n := fibonacci(3)
	fmt.Println("fibonacci", n)
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
