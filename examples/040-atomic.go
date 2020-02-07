package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Sum of all the elements of the array "v" located between the indices
// start (inclusive) and end (exclusive)
func Sum(v []int, start, end int) int {
	s := 0
	for start < end {
		s += v[start]
		start++
	}
	return s
}

func main() {
	parallelTasks := runtime.GOMAXPROCS(0)

	v := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
		12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}

	// experiment: this will never panic
	for true {
		wg := sync.WaitGroup{}
		wg.Add(parallelTasks)

		totalSum := int64(0)

		for t := 0; t < parallelTasks; t++ {
			s := t
			go func() {
				sum := Sum(v, s*len(v)/parallelTasks, (s+1)*len(v)/parallelTasks)

				atomic.AddInt64(&totalSum, int64(sum))

				wg.Done()
			}()
		}

		wg.Wait()
		if totalSum != 276 {
			panic(fmt.Sprint("Totalsum: ", totalSum))
		}

	}
	//fmt.Println("total sum: ", totalSum)
}

