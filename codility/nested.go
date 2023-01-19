package codility

import "fmt"

func Nested() {
	s := ""
	b := nested(s)
	fmt.Println(b)
}

func nested(S string) int {
	if len(S) == 0 {
		return 1
	}

	stack := []rune{}
	var pop rune

	for _, s := range S {
		switch s {
		case '(':
			stack = append(stack, s)
		case ')':
			if len(stack) == 0 {
				return 0
			}
			pop, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if pop != '(' {
				return 0
			}
		}
	}

	if len(stack) != 0 {
		return 0
	}

	return 1
}
