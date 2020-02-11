package main

import (
	"fmt"
	"time"
)

func ultraSlowTask() chan struct{} {
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		fmt.Println("Doing some ultra-slow task....")
		time.Sleep(5 * time.Second)
		fmt.Println("finished!")
	}()
	return ch
}

func main() {

	done := ultraSlowTask()
	fmt.Println("doing something in the meanwhile")
	<-done
}
