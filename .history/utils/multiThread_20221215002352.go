package utils

import (
	"fmt"
	"time"
)

func MultiThread() {
	multiThread()
}

func multiThread() {
	ch := make(chan time.Duration)
	ch2 := make(chan int)

	go hugeIter(ch)
	go hugeIter(ch)

	t := <-ch
	n := <- ch2

	fmt.Println("multi thread: ", t)
	fmt.Println("num: ", n)
}

func hugeIter(ch chan time.Duration) {
	start := time.Now()
	for i := 0; i < 100000; i++ {
	}
	elapse := time.Since(start)

	ch <- elapse
}

func hugeNum(ch chan time.Duration) {
	num := 0
	for i := 0; i < 100000; i++ {
		num += i
	}

	ch <-
}
