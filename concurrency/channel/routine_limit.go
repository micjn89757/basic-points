package channel

import (
	"fmt"
	"runtime"
	"time"
)

// 限制程序总协程数量，这里限制的是创建的另外协程

type GoroutineLimit struct {
	limit int
	ch chan struct{}
}

func NewGoroutineLimit(n int) *GoroutineLimit {
	return &GoroutineLimit{
		limit: n,
		ch: make(chan struct{}, n),
	}
}

func (g *GoroutineLimit) Run(f func()) {
	// 创建一个协程就增加计数
	g.ch <- struct{}{}
	go func() {
		f()

		// 执行完释放计数
		<- g.ch
	}()
}


func RoutineLimit() {
	// ticker.C是一个channel, ticker会根据时间间隔向其中填充，ticker会启动一个goroutine
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		// 每隔一秒打印一次goroutine数量
		for {
			<- ticker.C
			fmt.Printf("当前协程数量: %d\n", runtime.NumGoroutine())
		}
	}()

	limiter := NewGoroutineLimit(100)

	work := func() { // 函数
		// 处理逻辑
		time.Sleep(10 * time.Second)
	}

	for i := 0; i < 10000; i++ {
		limiter.Run(work)
	}
}