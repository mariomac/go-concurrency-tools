package main

import (
	"fmt"
	"sync"
)

type AverageStringLength struct{
	mutex sync.Mutex
	count int
	sum int
}

func (p *AverageStringLength) Account(sentence string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.count++
	p.sum += len(sentence)
}

func (p *AverageStringLength) PrintStats() {
	fmt.Printf("the average string length is %d (from %d strings)\n", p.sum/p.count, p.count)
}

func main() {
	const parallelTasks = 10

	// This program won't panic as long as we have used a mutex to perform accounting
	// operations in an atomic way
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
