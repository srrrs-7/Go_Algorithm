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
	ch := make(chan time.Duration)
	ch2 := make(chan int)

	hugeIter(ch)
	hugeNum(ch2)

	t := <-ch
	n := <-ch2

	fmt.Println("single thread: ", t)
	fmt.Println("single thread num: ", n)
}

func hugeIter(ch chan time.Duration) time.Duration {
	start := time.Now()
	for i := 0; i < 100000; i++ {
	}
	elapse := time.Since(start)
	ch <- elapse
	return elapse
}

func hugeNum(ch chan int) int {
	num := 0
	for i := 0; i <= 100000; i++ {
		num = i
	}
	ch <- num
	return num
}
