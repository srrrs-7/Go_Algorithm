package main

import (
	"github.com/sRRRs-7/Go_Algorithm.git/algorithm"
	atCoder "github.com/sRRRs-7/Go_Algorithm.git/atCorder"
	"github.com/sRRRs-7/Go_Algorithm.git/codility"
	"github.com/sRRRs-7/Go_Algorithm.git/leetcode"
	"github.com/sRRRs-7/Go_Algorithm.git/regex"
	analyzer "github.com/sRRRs-7/Go_Algorithm.git/staticAnalizer"
	"github.com/sRRRs-7/Go_Algorithm.git/test"
	"github.com/sRRRs-7/Go_Algorithm.git/utils"
)

const is_io = false
const is_leet = false

func main() {
	// leetcode
	leetcode.ReverseWords()
	leetcode.SelfCrossing()
	leetcode.AggregateComponent()
	leetcode.TotalAppealString()
	leetcode.GroupOfStrings()
	leetcode.CountIdealArray()
	leetcode.MinimumCostAve()
	leetcode.SubSequence()
	leetcode.MatchStr()
	leetcode.MaxDot()
	leetcode.LongestPrefix()
	leetcode.MKAve()
	//atCorder
	atCoder.Parenthesis()
	atCoder.GridMath()
	atCoder.Dp()
	if is_leet {
		atCoder.CutCake()
	}
	// utils
	utils.BaseNumberConvert()
	utils.ArrayToInteger()
	utils.StringContains()
	utils.SequenceArray()
	utils.MultiThread()
	utils.Contains()
	utils.Cipher()
	utils.ConvertBase64()
	utils.Aes()
	if is_io {
		utils.Io() // if you execute this function, is_io change true
	}
	// staticAnalyzer
	analyzer.StaticAnalyzer()
	// regex
	regex.RegexFunc()
	// language processing
	// mecab.NewMecab("こんにちは")
	// hebon := hebon.ToHebon("こんにちは")
	// fmt.Println(hebon)
	// algorithm
	algorithm.EuclidMethod()
	algorithm.MergeSort()
	algorithm.Distinct()
	algorithm.AddArray()
	algorithm.BinarySearch()
	algorithm.SortAlgo()
	algorithm.QuickSort()
	algorithm.MaxSubArray()
	algorithm.SequenceArray()
	algorithm.PrefixSum()
	algorithm.Sum()
	algorithm.BinaryTreeMain()
	algorithm.MaxMinArray()
	algorithm.Fibonacci()
	// codility
	codility.Demo()
	codility.BinaryGap()
	codility.CyclicRotation()
	codility.OddOccurrenceInArray()
	codility.FlogJump()
	codility.Permutation()
	codility.TapeArray()
	codility.FlogRiver()
	codility.PermCheck()
	codility.MaxCounter()
	codility.MissingInteger()
	codility.PassingCar()
	codility.CountDiv()
	codility.RangeQuery()
	codility.MinAveSlice()
	codility.Distinct()
	codility.MaxProductOfThree()
	codility.Triangle()
	codility.BucketStack()
	codility.GreetFish()
	codility.Nested()
	codility.StoneWall()
	codility.Dominator()
	codility.EqualReader()
	codility.MaxProfit()
	codility.MaxSliceSum()
	codility.Flags()
	// codility test
	test.Solution3()
	test.Solution4()
	test.Solution5()
	test.Solution6()
}
