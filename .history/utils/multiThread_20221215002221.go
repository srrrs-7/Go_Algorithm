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

	go hugeIter(ch)
	go hugeIter(ch)

	t := <-ch

	fmt.Println("multi thread: ", t)
}

func hugeIter(ch chan time.Duration) {
	start := time.Now()
	for i := 0; i < 100000; i++ {
	}
	elapse := time.Since(start)

	ch <- elapse
}
