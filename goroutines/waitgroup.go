package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Goroutine: ", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("foo")
		wg.Done()
	}()

	go func() {
		fmt.Println("bar")
		wg.Done()
	}()

	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Goroutine: ", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Goroutine: ", runtime.NumGoroutine())
	fmt.Println("Done!")
}
