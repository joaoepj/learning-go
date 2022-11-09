package main

import (
	"context"
	"log"
	"time"
)

func doSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)
	printCh := make(chan int)

	go doAnother(ctx, printCh)

	for i := 0; i < 5; i++ {
		printCh <- i
	}
	cancelCtx()

	time.Sleep(100 * time.Millisecond)

	log.Println("finished")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	log.Println("context's key value is: ", ctx.Value("myKey"))
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				log.Printf("error: %s\n", err)
			}
			log.Println("finished")
			return
		case i := <-printCh:
			log.Printf("%d\n", i)
		}
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.LstdFlags)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	doSomething(ctx)
}
