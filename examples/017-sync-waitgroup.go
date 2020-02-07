package main

import (
	"fmt"
	"sync"
)

func main() {
	const numTasks = 3

	wg := sync.WaitGroup{}
	wg.Add(numTasks)

	for i := 0; i < numTasks; i++ {
		numTask := i // "i" is in the shared scope for all the goroutines
		go func() {
			// emphasize that this does not guarantee the order of completion. E.g.:
			// Running parallel task 2
			// Running parallel task 0
			// Running parallel task 1
			fmt.Printf("Running parallel task %d\n", numTask)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("All the parallel tasks have ended. Exiting now")
}
