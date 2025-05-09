/*
select 使用案例：超时控制
*/
package concurrency

import (
	"context"
	"fmt"
	"time"
)

const (
	WorkUseTime = time.Second * 500
	TimeOut	= time.Second * 100
)


// 模拟一个耗时长的任务
func LongTimeWork() int {
	time.Sleep(WorkUseTime)
	return 888
}


// 方案1
func Handle1() int {
	deadline := make(chan struct{}, 1)
	workDone := make(chan int, 1)

	// 把要控制超时的函数放到一个协程里
	go func ()  {
		n := LongTimeWork()

		workDone <- n
	}()

	go func() {
		time.Sleep(TimeOut)
		deadline <- struct{}{}
		// close(deadline) 也可以使得select里的<- deadline解除阻塞
	}()

	select{
	case n := <- workDone:
		fmt.Println("LongTimeWork return")
		return n
	case <- deadline:
		fmt.Println("LongTimeWork timeout")
		return 0
	}
}



// 方案2，直接使用time.After
func Handle2() int {
	workDone := make(chan int, 1)
	go func() {
		n := LongTimeWork()
		workDone <- n
	}()

	select {
	case n := <- workDone:
		fmt.Println("LongTimeWork return")
		return n
	case <- time.After(TimeOut):
		fmt.Println("LongTimeWork timeout")
		return 0
	}
}


// 方案3，使用context withcancel
func Handle3() int {
	// 通过显示sleep再调用cancel来实现对函数的超时控制
	ctx, cancel := context.WithCancel(context.Background())

	workDone := make(chan int, 1)
	go func() {
		n := LongTimeWork()

		workDone <- n
	}()

	go func() {
		time.Sleep(TimeOut)
		cancel()	// 关闭ctx
	}()

	select {
	case n := <- workDone:
		fmt.Println("LongTimeWork return")
		return n
	case <- ctx.Done(): // ctx.Done是一个管道，调用了cancel就会关闭这个管道
		fmt.Println("LongTimeWork timeout")
		return 0
	}
}


func Handle4() int {
	ctx, cancel := context.WithTimeout(context.Background(), TimeOut) // 借助于带超时的context来实现对函数的超时控制
	defer cancel() // 良好习惯，函数退出前调用cancel()

	workDone := make(chan int, 1)
	go func() {
		n := LongTimeWork()
		workDone <- n
	}()

	select {
	case n := <- workDone:
		fmt.Println("LongTimeWork return")
		return n
	case <- ctx.Done(): // ctx.Done是一个管道，context超时或者调用了cancel就会关闭这个管道
		fmt.Println("LongTimeWork timeout")
		return 0	
	}
}