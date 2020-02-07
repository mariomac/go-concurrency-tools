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
	ch := make(chan int) //unbuffered channel

	go Receiver(ch) // Receiver and sender need to go in different goroutines
	Sender(ch)
}
