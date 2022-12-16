package utils

import (
	"fmt"
	"time"
)

func MultiThread() {
	multiThread()
	singleThread()
}

func multiThread() {
	ch := make(chan time.Duration)
	ch2 := make(chan int)

	go hugeIter(ch)
	go hugeNum(ch2)

	t := <-ch
	n := <-ch2

	fmt.Println("multi thread: ", t)
	fmt.Println("multi thread num: ", n)
}

func singleThread() {
	t := HugeIter()
	n := HugeNum()

	fmt.Println("single thread: ", t)
	fmt.Println("single thread num: ", n)
}

// multi thread
func hugeIter(ch chan time.Duration) {
	start := time.Now()
	for i := 0; i < 1000000000; i++ {
	}
	elapse := time.Since(start)
	ch <- elapse
}

func hugeNum(ch chan int) {
	num := 0
	for i := 0; i <= 1000000000; i++ {
		num = i
	}
	ch <- num
}

// single thread
func HugeIter() time.Duration {
	start := time.Now()
	for i := 0; i < 1000000000; i++ {
	}
	elapse := time.Since(start)
	return elapse
}

func HugeNum() int {
	num := 0
	for i := 0; i <= 1000000000; i++ {
		num = i
	}
	return num
}
