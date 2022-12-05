package regex

import (
	"fmt"
	"regexp"
)

func RegexFunc() {
	fmt.Println(checkRegex("a*c", "aaadbccc"))
}

func checkRegex(reg, s string) bool {
	return regexp.MustCompile(reg).Match([]byte(s))
}