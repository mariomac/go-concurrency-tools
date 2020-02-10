package main

import (
	"fmt"
	"runtime"
	"sync"
)

type AverageStringLength struct{
	count int
	sum int
}

func (p *AverageStringLength) Account(sentence string) {
	p.count++
	p.sum += len(sentence)
}

func (p *AverageStringLength) PrintStats() {
	fmt.Printf("the average string length is %d (from %d strings)\n", p.sum/p.count, p.count)
}

func main() {
	runtime.GOMAXPROCS(4)
	const parallelTasks = 10

	wg := sync.WaitGroup{}
	wg.Add(parallelTasks)

	p := AverageStringLength{}

	for i := 0 ; i < parallelTasks ; i++ {
		go func() {
			defer wg.Done()
			p.Account("hello")
		}()
	}

	wg.Wait()
	p.PrintStats()

}
