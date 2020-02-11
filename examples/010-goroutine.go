package main

import (
	"fmt"
	"sync"
)

func main() {
	const numTasks = 5
	wg := sync.WaitGroup{}
	for i := 0; i < numTasks; i++ {
		numTask := i // "i" is in the shared scope for all the goroutines
		wg.Add(1)
		go func() {
			fmt.Printf("Running parallel task %d\n", numTask)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("All the parallel tasks have ended. Exiting now")
}
