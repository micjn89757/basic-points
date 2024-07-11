package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var n int32 = 1

func Addn() {
	// n++
	// 并发安全
	atomic.AddInt32(&n, 1)
}

func atomicRWlock() {
	// 开启1000个goroutine
	const P = 1000
	wg := sync.WaitGroup{}

	wg.Add(1000)
	for i := 0; i < P; i++ {
		go func ()  {
			Addn()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(n)
}