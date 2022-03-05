package main

import "fmt"

func main() {
	ch := make(chan string, 1)

	ch <- "Rizqi"

	fmt.Println(<-ch)
}
