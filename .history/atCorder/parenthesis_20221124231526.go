package atCorder

import "fmt"

func Parenthesis() {
	brackets := "(()())(())"

	res := parenthesis(brackets)
	fmt.Println(res)
}

func parenthesis(brackets string) bool {
	for _, b := range brackets {
		fmt.Println(b)
	}

	return true
}