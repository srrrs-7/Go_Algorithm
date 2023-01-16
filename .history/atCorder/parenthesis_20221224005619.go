package atCorder

import "fmt"

func Parenthesis() {
	brackets := "(()))"

	res := parenthesis(brackets)
	fmt.Println("parenthesis", res)
}

func parenthesis(brackets string) bool {
	cnt := 0

	for _, b := range brackets {
		switch string(b) {
		case "(":
			cnt += 1
		case ")":
			if cnt <= 0 {
				return false
			}
			cnt -= 1
		default:
			return false
		}
	}
	if cnt != 0 {
		return false
	} else {
		return true
	}
}
