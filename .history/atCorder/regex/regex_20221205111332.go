package regex

import (
	"fmt"
	"regexp"
)

func RegexFunc() {
	fmt.Println("regex check: ", checkRegex("a*c", "abcs"))
	fmt.Println("regex check: ", checkRegex("a?c", "abcs"))
	fmt.Println("regex check: ", checkRegex("a+c", "accccc"))
}

func checkRegex(reg, s string) bool {
	return regexp.MustCompile(reg).Match([]byte(s))
}