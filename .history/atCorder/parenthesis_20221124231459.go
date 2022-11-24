package atCorder

import "fmt"

func Parenthesis() {
	brackets := "(()())(())"

	res := parethesis(brackets)
	fmt.Println(res)
}

func parenthesis(brackets string) bool {
	for _, b := range brackets {
		fmt.Println(b)
	}
}