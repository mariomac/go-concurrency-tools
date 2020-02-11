package main

import (
	"fmt"
	"time"
)

func sweetDealer() <-chan string {
	ch := make(chan string, 10)
	go func() {
		defer close(ch)
		ch <- "Donut"
		//time.Sleep(5 * time.Second)
		ch <- "Cookie"
		//time.Sleep(5 * time.Second)
		ch <- "Sweet"
	}()
	return ch
}

func main() {

	sweetsCh := sweetDealer()

	time.Sleep(5 * time.Second)

	for sweet := range sweetsCh {
		fmt.Println("received a delicious", sweet)
	}
}
