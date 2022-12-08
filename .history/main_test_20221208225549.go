package main

import (
	"testing"

	"github.com/sRRRs-7/Go_Algorithm.git/leetcode"
)

func BenchmarkReverseWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.ReverseWords()
	}
}

func BenchmarkAggregateComponent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.AggregateComponent()
	}
}

func BenchmarkAggregateComponent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.AggregateComponent()
	}
}
