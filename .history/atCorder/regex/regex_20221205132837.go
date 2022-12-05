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
	fmt.Println("regex check: ", checkRegex(`^0[789]0-\d{4}-\d{4}$`, "080-3212-3456"))	// phone number regex
	fmt.Println("regex check: ", checkRegex(`^\d{4}:\d{2}:\d{2}$`, "1996:08:25"))	// date regex
	fmt.Println("regex check: ", checkRegex(`^https?://([\w-]+\.)+[\w-]+(/[\w-/?%&=]*)?$`, "https://amazon.com/asd/?%=sad"))	// URL regex
	fmt.Println("regex check: ", checkRegex(`^\d{3}-\d{4}$`, "035-2345"))	// post code

	// string retrieve
	fmt.Println("regex retrieve: ", regexRetriever(`[\d\-]+`, "hoge:1234-1234-1234 asda:1234-2341-1342"))	// find all regex

	result := regexRetriever(`(\S+):([\d\-]+)`, "hoge:1234-1234-1234 asda:1234-2341-1342")
	fmt.Println("regex retrieve: ", result[0], result[1])	// find all regex
}

func checkRegex(reg, s string) bool {
	return regexp.MustCompile(reg).Match([]byte(s))
}

func regexRetriever(reg, s string) [][]string {
	return regexp.MustCompile(reg).FindAllStringSubmatch(s, -1)
}