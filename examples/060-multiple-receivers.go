package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func sweetsDealer() <-chan string {
	ch := make(chan string, 10)
	go func() {
		defer close(ch)
		ch <- "Donut"
		time.Sleep(time.Second)
		ch <- "Cookie"
		time.Sleep(time.Second)
 		ch <- "Sweet"
	}()
	return ch
}

func sweetsEater(name string, sweetsCh <-chan string) {
	for sweet := range sweetsCh {
		fmt.Println(name, "will eat a delicious", sweet)
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	sweetsCh := sweetsDealer()
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		sweetsEater("Maria", sweetsCh)
		wg.Done()
	}()
	go func() {
		sweetsEater("Marta", sweetsCh)
		wg.Done()
	}()
	go func() {
		sweetsEater("Jose", sweetsCh)
		wg.Done()
	}()

	wg.Wait()
}
