package concurrency

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

const P = 6

func TestMapCmp(t *testing.T) {
	var wg sync.WaitGroup
 	
	runtime.GOMAXPROCS(2)
	// 6个协程, 写100次, 读100次
	// cmp := NewConcurrentHashMap[int](6, 600)
	cmps := NewConcurrentHashMapSpin[int](6, 600)
	begin := time.Now()
	wg.Add(6)
	for i := 0; i < 6; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000000; i++ {
				cmps.SetS(i, i)
			}

			for i := 0; i < 1000000; i++ {
				cmps.GetS(i)
			}
		}()
	}
	wg.Wait()

	totalTime := time.Since(begin).Milliseconds()

	t.Log(totalTime)
}