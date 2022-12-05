package regex

import "regexp"

func RegexFunc() {
	checkIter("")
}

func checkIter(reg, s string) bool {
	return regexp.MustCompile(reg).MatchString(s)

}