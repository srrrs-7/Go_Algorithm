package main

import (
	"github.com/sRRRs-7/Go_Algorithm.git/algorithm"
	"github.com/sRRRs-7/Go_Algorithm.git/colly"
	"github.com/sRRRs-7/Go_Algorithm.git/leetcode"
	analyzer "github.com/sRRRs-7/Go_Algorithm.git/staticAnalizer"
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
	leetcode.AggregateComponent()
	// staticAnalyzer
	analyzer.StaticAnalyzer()
}