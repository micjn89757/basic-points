/*
排查协程泄露
*/
package concurrency

import (
	"context"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	WorkTime = time.Millisecond * 5
)

// 客户端访问100000次
func client() {
	for i := 0; i < 100000; i++ {
		http.Get("http://127.0.0.1:3000")
	}
}

func rpc() int {
	time.Sleep(WorkTime)
	return 888
}

func homeHandler(ctx *gin.Context) {
	TimeoutCtx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	workDone := make(chan int, 1)	// 如果是非阻塞
	go func() {
		n := rpc()
		workDone <- n
	}()

	select {
	case n := <- workDone:
		ctx.String(http.StatusOK, strconv.Itoa(n))
	case <- TimeoutCtx.Done():
		ctx.String(http.StatusInternalServerError, strconv.Itoa(0))
	}
}


func server() {
	ticker := time.NewTicker(1 * time.Second) 
	defer ticker.Stop() 
	
	go func() { // 每隔一秒打印一次协程数量 
		for {
			<- ticker.C
			fmt.Printf("当前协程数: %d\n", runtime.NumGoroutine())
		}
	}()

	go http.ListenAndServe("127.0.0.1:3000", nil)

	gin.DefaultWriter = io.Discard
	engine := gin.Default()
	engine.GET("/", homeHandler)
	engine.Run("127.0.0.1:3000")
}