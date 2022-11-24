package atCorder

import "fmt"

func Parenthesis() {
	brackets := "(()())(())"

	res := parenthesis(brackets)
	fmt.Println("parenthesis: ", res)
}

func parenthesis(brackets string) bool {
	arr := []string{}

	for _, b := range brackets {
		switch string(b) {
		case "(" :
			arr = append(arr, string(b))
		case ")":
			if len(arr) <= 1 { return false }
			arr = arr[0 : len(arr)-2]
		default:
			continue
		}
	}

	return true
}