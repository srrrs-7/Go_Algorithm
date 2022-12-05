package regex

import (
	"fmt"
	"regexp"
)

func RegexFunc() {
	fmt.Println("regex check: ", checkRegex("a*c", "c"))
	fmt.Println("regex check: ", checkRegex("a?c", "a"))
	fmt.Println("regex check: ", checkRegex("a+c", "a"))
}

func checkRegex(reg, s string) bool {
	return regexp.MustCompile(reg).Match([]byte(s))
}