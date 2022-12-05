package regex

import "regexp"

func RegexFunc() {
	checkIter("")
}

func checkIter(reg, s string) bool {
	b := regexp.MustCompile(reg).MatchString(s)
}