package main

import "fmt"

func main() {
	a := 0
	for a <= 5 {
		fmt.Println(a)
		a++
	}

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}
