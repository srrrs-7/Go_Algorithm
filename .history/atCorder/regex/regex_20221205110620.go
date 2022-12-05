package regex

import "regexp"

func RegexFunc() {
	checkIter("a*c", "aaac")
}

func checkIter(reg, s string) bool {
	return regexp.MustCompile(reg).MatchString(s)

}