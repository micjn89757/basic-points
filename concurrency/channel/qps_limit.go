package channel

import (
	"sync"
	"time"
)

// 限制接口的并发请求，瞬时最高并发请求量
var qps = make(chan struct{}, 100)


func handler() {
	qps <- struct{}{}
	defer func ()  {
		<- qps
	}()

	// 业务处理
	time.Sleep(3 * time.Second)
}

func limitConcurr() {
	const P = 100
	wg := sync.WaitGroup{}
	wg.Add(P)

	for i := 0; i < P; i++ {
		go func (){
			defer wg.Done()
			handler()
		}()
	}

	wg.Wait()
}