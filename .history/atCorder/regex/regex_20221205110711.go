package regex

import (
	"fmt"
	"regexp"
)

func RegexFunc() {
	fmt.Println(checkIter("a*c", "aaac"))
}

func checkIter(reg, s string) bool {
	return regexp.MustCompile(reg).MatchString(s)
}