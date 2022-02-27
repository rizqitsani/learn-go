package main

import "fmt"

func factorial(i int) int {
	if i == 0 || i == 1 {
		return 1
	}

	return i * factorial(i-1)
}

func main() {
	fmt.Println(factorial(1))
	fmt.Println(factorial(5))
}
