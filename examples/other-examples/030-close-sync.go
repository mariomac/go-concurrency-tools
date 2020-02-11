package main

import "fmt"

func SyncedTask() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		fmt.Println("doing something in parallel...")
		for i := 0; i < 3; i++ {
			fmt.Println(i, "...")
		}
		fmt.Println("finished parallel task")

		close(ch) // any receive operation will be unblocked
	}()
	return ch
}

func main() {

	fmt.Println("Running some task in parallel...")

	wait := SyncedTask()

	<-wait // channel output will be ignored

	fmt.Println("program exit")
}
