package main

import (
	"context"
	"fmt"
	"time"
)

// the way of cancelling workloads in Go is by means of contexts
func cancellableFunction(ctx context.Context, ch chan<- string) {
	for {
		select {
		case <-time.After(2 * time.Second):
			// Running this every 5 seconds, unless the context is cancelled
			ch <- "ping!"
		case <-ctx.Done():
			fmt.Println("Function has been cancelled. Stopped doing stuff")
			return
		}
	}
}

func main() {
	ch := make(chan string, 10)
	done := make(chan struct{})
	// we create a cancellable context from the background context
	ctx, doCancel := context.WithCancel(context.Background())

	go func() {
		cancellableFunction(ctx, ch)
		close(done)
	}()

	fmt.Println("received", <-ch)
	fmt.Println("received", <-ch)

	fmt.Println("Tired of this, just cancelling the background task and waiting for it to be finished")
	doCancel()
	<-done

	fmt.Println("exiting")
}
