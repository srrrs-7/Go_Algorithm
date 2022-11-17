package colly

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Data struct {
	data []string
}

func Scraping() {
	c := colly.NewCollector()
	fmt.Println(c)

}