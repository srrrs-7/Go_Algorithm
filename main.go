package main

import (
	"github.com/sRRRs-7/Go_Algorithm.git/algorithm"
	"github.com/sRRRs-7/Go_Algorithm.git/colly"
	"github.com/sRRRs-7/Go_Algorithm.git/leetcode"
	staticanalizer "github.com/sRRRs-7/Go_Algorithm.git/static_analizer"
)

func main() {
	colly.Scraping()
	// algorithm
	algorithm.BinarySearch()
	algorithm.EuclidMethod()
	algorithm.MergeSort()
	// leetcode
	leetcode.ReverseWords()
	leetcode.SelfCrossing()
	// staticAnalyzer
	staticanalizer.StaticAnalyzer()
}