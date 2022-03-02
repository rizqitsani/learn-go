package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Goroutine: ", runtime.NumGoroutine())

	var counter int64 = 0
	numberOfLoops := 50

	var wg sync.WaitGroup
	wg.Add(numberOfLoops)

	for i := 0; i < numberOfLoops; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter: ", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Goroutine: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
