package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Goroutine: ", runtime.NumGoroutine())

	counter := 0
	numberOfLoops := 50

	for i := 0; i < numberOfLoops; i++ {
		go func() {
			newCounter := counter
			runtime.Gosched()
			newCounter++
			counter = newCounter
			fmt.Println("Counter: ", counter)
		}()
	}

	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Goroutine: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
