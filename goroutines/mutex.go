package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Goroutine: ", runtime.NumGoroutine())

	counter := 0
	numberOfLoops := 50

	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(numberOfLoops)

	for i := 0; i < numberOfLoops; i++ {
		go func() {
			mutex.Lock()
			counter++
			fmt.Println("Counter: ", counter)
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Goroutine: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
