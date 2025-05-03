package concurrency

import (
	"context"
	"log"
	"time"
	// "sync"
	// "time"
)

// var wg sync.WaitGroup

type TraceCode string


func wVPWorker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在goroutine中获取trace code

	if !ok {
		log.Printf("invalid trace, code")
	}
LOOP:
	for {
		log.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond*10) // 假设正常连接数据库10ms
		select {
		case <- ctx.Done(): // 50ms后自动调用
			break LOOP 
		default:
		}
	}
	log.Printf("worker done!")
	wg.Done()
}

func withValuePractice() {
	// 设置一个50ms的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	wg.Add(1)
	go wVPWorker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知goroutine结束
	wg.Wait()
	log.Printf("over")
}

func wTPWorker(ctx context.Context) {
LOOP:
	for {
		log.Printf("db connecting ...")
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <- ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	log.Printf("worker done!")
	wg.Done()
}

func withTimeoutPractice() {
	ctx, cancel := context.WithTimeout(context.Background(), 50 * time.Millisecond)
	wg.Add(1)
	go wTPWorker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知goroutine结束
	wg.Wait()
	log.Printf("over")
}

func withDeadlinePractice() {
	d := time.Now().Add(50 * time.Millisecond)

	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <- time.After(1 * time.Second):
		log.Printf("overslept")
	case <- ctx.Done():
		log.Printf("%v\n", ctx.Err())
	}
}


func wCPGen(ctx context.Context) <- chan int {
	dst := make(chan int)
	n := 1

	go func()  {
		for {
			log.Printf("goroutine")
			select {
			case <- ctx.Done():
				return // return结束goroutine防止泄露
			case dst <- n:
				n++
			}
		}
	}()

	return dst
}

func withCancelPractice() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 获取完需要的整数后调用cancel

	for n := range wCPGen(ctx) {
		log.Printf("%d\n", n)

		if n == 5 {
			break
		}
	}
}

// // context官方示例
// var wg sync.WaitGroup

// func worker(ctx context.Context) {
// LOOP:
// 	for {
// 		log.Printf("worker...")
// 		time.Sleep(time.Second)
// 		select {
// 		case <- ctx.Done(): // 等待上级通知
// 			break LOOP
// 		default:
// 		}
// 	}

// 	wg.Done()
// }

// func mainGorou() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	wg.Add(1)
// 	go worker(ctx)
// 	time.Sleep(time.Second * 3)
// 	cancel() // 通知goroutine结束
// 	wg.Wait()
// 	log.Printf("over")
// }