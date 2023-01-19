package codility

import (
	"fmt"
)

func BucketStack() {
	s := "]"
	b := bucket(s)
	fmt.Println(b)
}

func bucket(S string) int {
	stack := make([]rune, 0)
	var pop rune

	for _, v := range S {
		switch {
		case v == '(':
			stack = append(stack, v)
		case v == '[':
			stack = append(stack, v)
		case v == '{':
			stack = append(stack, v)
		case v == ')':
			if len(stack) == 0 {
				return 0
			}
			pop, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if pop != '(' {
				return 0
			}
		case v == ']':
			if len(stack) == 0 {
				return 0
			}
			pop, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if pop != '[' {
				return 0
			}
		case v == '}':
			if len(stack) == 0 {
				return 0
			}
			pop, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if pop != '{' {
				return 0
			}
		}
	}
	if len(stack) != 0 {
		return 0
	}
	return 1
}
