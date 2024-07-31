package context

import (
	"context"
	"fmt"
)

func step1(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "name", "step1")
	return child
}

func step2(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "age", 18)

	return child
}


func step3(ctx context.Context) {
	fmt.Printf("1.%s\n", ctx.Value("name"))
	fmt.Printf("2.%d\n", ctx.Value("age"))
}

