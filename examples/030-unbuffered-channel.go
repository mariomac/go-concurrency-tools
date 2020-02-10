package main

import "fmt"

type Ball struct{}

func main() {
	ch := make(chan Ball)

	b := Ball{}
	fmt.Printf("The pitcher throws %T\n", b)
	ch <- b //deadlock

	rcv := <-ch
	fmt.Printf("The catcher received %T\n", rcv)
}
