package utils

import (
	"fmt"
	"math"
)

func BaseNumberConvert() {
	num := 10
	base := 2

	res := baseNumberConvert(num, base)
	fmt.Println("convert base number", res)
}

func baseNumberConvert(num, base int) int {
	remainder := 0
	quotient := 0
	result := 0
	cnt := 1.0

	for quotient != 1 {
		remainder = num % base
		quotient = num / base
		num = quotient
		result = remainder * (int(math.Pow(10, cnt)))
		cnt += 1.0
		fmt.Println(math.Pow(10, cnt))
	}

	return result
}