package main

import (
	"fmt"
	"runtime"
)

type Ball struct{}

func main() {
	runtime.GOMAXPROCS(4)
	ch := make(chan Ball, 10)

	b := Ball{}

	fmt.Printf("The pitcher throws %T\n", b)
	ch <- b

	rcv := <-ch
	fmt.Printf("The catcher received %T\n", rcv)

}
