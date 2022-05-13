package main

import (
	"context"
	"fmt"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request_id", "my request id")
}

func doSmthCool(ctx context.Context) {
	fmt.Println(ctx.Value("request_id"))
}

func main() {
	ctx := context.Background()
	ctx = enrichContext(ctx)
	doSmthCool(ctx)
}
