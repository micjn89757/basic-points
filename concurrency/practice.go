package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// 全局等待组遍历
var wg sync.WaitGroup

// 单向通道
// <- chan int 只接手通道
// chan <- int 只发送通道

// len获取通道中元素数量，cap获取容量
// 无缓冲channel,同步channel
func nioChannel() {
	ch := make(chan int)
	defer close(ch)
	go func (ch chan int)  {
		ret := <- ch
		fmt.Println(ret)
	}(ch)

	ch <- 10
	fmt.Println("发送成功")
}


// 有缓冲channel
func bioChannel() {
	ch := make(chan int, 1)
	
	ch <- 10 // 不需要接收者不会报错
	fmt.Println("发送成功")
}

// 遍历channel
func traveseChannel() {
	ch := make(chan int, 3)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)

		for i := 0; i < 10; i++ {
			ch <- i
		}

		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	wg.Add(1)
	// 遍历并取走管道中的元素
	go func() {
		defer wg.Done()
		for e := range ch {  // 通道关闭后，会在通道内所有值被接收完毕后退出循环
			fmt.Println(e)
		}
		

		// 接收方式2
		// for {
		// 	e, ok := <- ch  // 通道关闭返回false
		// 	if ok {
		// 		fmt.Println(e)
		// 	} else {
		// 		break
		// 	}
		// }
		fmt.Println("bye")
	}()

	wg.Wait()
}


// channel传递信号，通常使用空结构体(不占内存，取地址会返回统一的值)，表示不传输数据，只传递信号
func signalChannel() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("子协程结束")
		ch <- struct{}{}
	}()

	<- ch
}