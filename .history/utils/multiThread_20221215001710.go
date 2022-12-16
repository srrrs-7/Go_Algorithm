package utils

import "time"

func MultiThread() {
	multiThread()
}

func multiThread() {
	go hugeIter()
}

func hugeIter() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
	}
	elapse := time.Since(start)
}
