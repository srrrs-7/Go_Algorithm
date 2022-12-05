package regex

import (
	"fmt"
	"regexp"
)

func RegexFunc() {
	fmt.Println("regec check", checkRegex("a*c", "aaacd"))
}

func checkRegex(reg, s string) bool {
	return regexp.MustCompile(reg).Match([]byte(s))
}