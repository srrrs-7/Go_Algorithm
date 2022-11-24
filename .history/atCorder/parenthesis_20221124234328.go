package atCorder

import "fmt"

func Parenthesis() {
	brackets := "(()())(())"

	res := parenthesis(brackets)
	fmt.Println("parenthesis", res)
}

func parenthesis(brackets string) bool {
	arr := []string{}

	for i, b := range brackets {
		switch string(b) {
		case "(" :
			arr = append(arr, string(b))
		case ")":
			arr = arr[0 : i-2]
		default:
			continue
		}
	}

	return true
}