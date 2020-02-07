package main

import "fmt"

func printFiveTimes(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("(%d of 5) %s\n", i+1, msg)
	}
}

func main() {
	go printFiveTimes("This goroutine may not always end")
	printFiveTimes("This will be printed exactly five times")

	fmt.Println("Finishing program")
}
