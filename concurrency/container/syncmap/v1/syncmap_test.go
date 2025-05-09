package v1

import (
	"runtime"
	"sync"
	"testing"
)

const G = 6

func BenchmarkSyncMap(b *testing.B) {
	var wg sync.WaitGroup

	runtime.GOMAXPROCS(4)

	// 6 goroutines, write 10000, read 10000
	syncmap := NewSyncMap[int, int](1)
	b.ResetTimer()
	wg.Add(G)
	for range G {
		go func() {
			defer wg.Done()
			for i := range b.N {
				syncmap.Set(i, i)
			}

			for i := range b.N {
				syncmap.Get(i)
			}
		}()
	}

	wg.Wait()
}
