package main

import (
	"errors"
	"fmt"
)

// pseudo completable future
func Operation() (<-chan int, <-chan error) {
	resultCh := make(chan int)
	errCh := make(chan error)

	go func() {
		errCh <- errors.New("something wrong happened")
	}()

	return resultCh, errCh
}

func main() {

	resultCh, errCh := Operation()

	select {
	case ret := <-resultCh:
		fmt.Println("the operation completed with a returned value: ", ret)
	case err := <-errCh:
		fmt.Println("the operation failed:", err.Error())
	}

}
