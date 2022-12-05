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
	fmt.Println("regex check: ", checkRegex(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, "sahdfldk@alskdjfh.fsd"))	// email regex
	fmt.Println("regex check: ", checkRegex(`^0[789]0-\d{4}-\d{4}$`, "040-3212-3456"))	// email regex

}

func checkRegex(reg, s string) bool {
	return regexp.MustCompile(reg).Match([]byte(s))
}