package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	const stopAtMB = 324
	for {
		wait := make(chan struct{})
		memstats := runtime.MemStats{}
		go func() {
			runtime.ReadMemStats(&memstats)
			mb := memstats.Sys/1_000_000
			fmt.Println("active goroutines =", runtime.NumGoroutine(),
				" (total memory:", mb , "MB)")
			if mb >= stopAtMB {
				os.Exit(0)
			}
			close(wait)
			time.Sleep(100 * time.Hour)
		}()
		<-wait
	}
}
