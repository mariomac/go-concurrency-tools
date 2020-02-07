package main

import "fmt"

func Sender(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Printf("%d successfully sent\n", i)
	}
}

func Receiver(ch <-chan int) {
	for {
		num := <-ch
		fmt.Printf("received: %d\n", num)
	}
}

func main() {
	ch := make(chan int, 10) //buffered channel

	// With a buffered channel, receiver and sender don't strictly need to go in
	// different goroutines
	go Receiver(ch)
	Sender(ch)
}
