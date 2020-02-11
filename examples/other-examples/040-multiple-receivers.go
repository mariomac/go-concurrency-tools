package main

import (
	"fmt"
	"time"
)

func Eater(name string, sweets <-chan string) {
	for {
		sweet := <-sweets
		fmt.Println(name, "will eat", sweet)
	}
}

func main() {
	sweets := make(chan string)

	for _, name := range []string{"Mark", "Aina", "Judit"} {
		go Eater(name, sweets)
	}

	for _, s := range []string{"Donut", "Croissant", "Ensaimada"} {
		// There is no guarantee of a fair delivery
		sweets <- s
	}

	// give room to allow all the eaters printing the message
	time.Sleep(2 * time.Second)
}
