package main

import (
	"github.com/sRRRs-7/Go_Algorithm.git/algorithm"
	"github.com/sRRRs-7/Go_Algorithm.git/colly"
	"github.com/sRRRs-7/Go_Algorithm.git/leetcode"
)

func main() {
	colly.Scraping()
	// algorithm
	algorithm.BinarySearch()
	// leetcode
	leetcode.ReverseWords()
	leetcode.SelfCrossing()
}