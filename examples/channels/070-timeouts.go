package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("Calculating the Answer to the Ultimate Question of Life, the Universe, and Everything...")
		time.Sleep(15 * time.Second)
		ch <- 42
	}()

	fmt.Println("Waiting for a response...")
	select {
	case ret := <-ch:
		fmt.Println("Received:", ret)
	case <-time.After(2 * time.Second):
		fmt.Println("Error: timeout while waiting for a response")
	}

}
