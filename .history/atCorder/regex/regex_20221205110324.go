package regex

import "regexp"

func RegexFunc() {
	r := regexp.MustCompile("abc")
	checkIter(r)
}

func checkIter(r *regexp.Regexp) {}