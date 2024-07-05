package demo

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 全局等待组变量
var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func() {
		defer wg.Done()
		go func() {
			defer wg.Done()
			fmt.Println("child")
		}()

		for i := 'a'; i < 'z'; i++ {
			fmt.Printf("%d\n", i)
		}
	}()

	fmt.Println("main") // 表示main协程

	wg.Wait() // 阻塞，直到计数减为0

	cpuN := runtime.NumCPU() // 获取CPU逻辑核心数

	fmt.Println("逻辑核数", cpuN)

	runtime.GOMAXPROCS(cpuN / 2) // 限制go进程最多使用的核心数

	const P = 100000
	for i := 0; i < P; i++ {
		go time.Sleep(time.Second * 10)
	}
	fmt.Println("进程中存活的协程数", runtime.NumGoroutine())

}
