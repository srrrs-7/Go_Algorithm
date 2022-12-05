package regex

import (
	"fmt"
	"regexp"
)

func RegexFunc() {
	fmt.Println("regex check: ", checkRegex(`a*c`, "c"))
	fmt.Println("regex check: ", checkRegex(`a?c`, "c"))
	fmt.Println("regex check: ", checkRegex(`a+c`, "ac"))

	fmt.Println("regex check: ", checkRegex(`[a-z]`, "g"))
	fmt.Println("regex check: ", checkRegex(`[^a-z]`, "7"))
	fmt.Println("regex check: ", checkRegex(`[0-9]`, "7"))
	fmt.Println("regex check: ", checkRegex(`[^0-9]`, "g"))

	// useful
	fmt.Println("regex check: ", checkRegex(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, "sahdfldk@alskdjfh"))
}

func checkRegex(reg, s string) bool {
	return regexp.MustCompile(reg).Match([]byte(s))
}