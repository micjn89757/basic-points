package gocontext

import (
	"context"
)

// WithCancel示例
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()

	return dst
}
