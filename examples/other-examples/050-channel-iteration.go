package main

import (
	"fmt"
	"sync"
)

func Eater(name string, sweets <-chan string) {
	for sweet := range sweets {
		fmt.Println(name, "will eat", sweet)
	}
	fmt.Println("(", name, ") No more sweets? bye")
}

func main() {
	sweets := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(3)

	for _, name := range []string{"Mark", "Aina", "Judit"} {
		n := name
		go func() {
			defer wg.Done()
			Eater(n, sweets)
		}()
	}

	for _, s := range []string{"Donut", "Croissant", "Ensaimada"} {
		sweets <- s
	}
	close(sweets) // Important! or Eater goroutines won't end

	// don't need to sleep... just wait for completions
	wg.Wait()
	fmt.Println("All the sweets were delivered and eaten")
}
