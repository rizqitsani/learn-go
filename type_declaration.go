package main

import "fmt"

type CustomInt int

func main() {
	var age CustomInt
	age = 19

	fmt.Printf("%v %T\n", age, age)
	// %v -> value
	// %T -> type

	var hour = 23
	fmt.Printf("%v %T\n", hour, hour)
}
