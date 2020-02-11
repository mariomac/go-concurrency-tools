package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AverageStringLength struct{
	count int64
	sum int64
}

func (p *AverageStringLength) Account(sentence string) {
	atomic.AddInt64(&p.count, 1)
	atomic.AddInt64(&p.sum, int64(len(sentence)))
}

func (p *AverageStringLength) PrintStats() {
	fmt.Printf("the average string length is %d (from %d strings)\n", p.sum/p.count, p.count)
}

func main() {
	const parallelTasks = 10

	// Like the previous program, but using the atomic package instead of a mutex
	for {
		wg := sync.WaitGroup{}
		wg.Add(parallelTasks)

		p := AverageStringLength{}

		for i := 0; i < parallelTasks; i++ {
			go func() {
				defer wg.Done()
				p.Account("hello")
			}()
		}

		wg.Wait()
		if p.count != 10 || p.sum != 50 {
			p.PrintStats()
			panic("pum!")
		}
	}

}
