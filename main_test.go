package main

import (
	"testing"

	"github.com/sRRRs-7/Go_Algorithm.git/algorithm"
)

func BenchmarkReverseWords(b *testing.B) {
	for i:=0; i<b.N; i++ {
		algorithm.ReverseWords()
	}
}