package main

import (
	"fmt"
)

func main() {
	const numTasks = 3

	for i := 0; i < numTasks; i++ {
		numTask := i // "i" is in the shared scope for all the goroutines
		go func() {
			fmt.Printf("Running parallel task %d\n", numTask)
		}()
	}
	fmt.Println("All the parallel tasks have ended. Exiting now")
}
