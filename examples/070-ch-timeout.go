package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func sweetssDealer() <-chan string {
	ch := make(chan string, 10)
	go func() {
		defer close(ch)
		time.Sleep(10 * time.Second)
		ch <- "Donut"
	}()
	return ch
}

func main() {
	runtime.GOMAXPROCS(4)

	sweetsCh := sweetssDealer()

	select {
	case sw := <-sweetsCh:
		fmt.Println("receiving", sw)
	case <-time.After(5 * time.Second):
		fmt.Println("timeout while waiting for a sweet")
	}


	ctx.Done()
}
