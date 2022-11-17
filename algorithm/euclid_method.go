package algorithm

import "fmt"

func euclidMethod(num1, num2 int) (gcd int) {
	for {
		remainder:= num1 % num2
		if remainder == 0 {
			return num2
		}
		num1 = num2
		num2 = remainder
	}
}

func EuclidMethod() {
	num1 := 10
	num2 := 25
	res := euclidMethod(num1, num2)
	fmt.Println("euclid method: ", res)
}