package gocontext

import (
	"context"
	"testing"
)

func TestWithCacnel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // control goroutine

	for n := range gen(ctx) {
		t.Log(n)
		if n == 5 {
			break
		}
	}
}
