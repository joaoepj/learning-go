package main

import (
	"fmt"
	"context"
)

func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: context's key is %s\n", ctx.Value("myKey"))
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	doSomething(ctx)
}