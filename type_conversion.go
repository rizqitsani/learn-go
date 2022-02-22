package main

import "fmt"

type CustomInt int

func main() {
	var age CustomInt
	age = 19

	fmt.Println(age)

	var hour = 23 // int

	age = CustomInt(hour)
	fmt.Println(age)
}
