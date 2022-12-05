package regex

import "regexp"

func RegexFunc() {
	r := regexp.MustCompile("abc")
}

func checkIter(r *regexp.Regexp) {}