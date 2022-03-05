package main

import "fmt"

func main() {
	channel := generateNumbers()

	receive(channel)
}

func generateNumbers() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func receive(ch <-chan int) {
	for number := range ch {
		fmt.Println(number)
	}
}
