package main

import "fmt"

type Animal struct {
	name string
}

func main() {
	var anyType interface{}

	anyType = 2
	fmt.Println(anyType)

	anyType = Animal{"Sappy"}
	fmt.Println(anyType)

	anyType = []int{1, 2, 3}
	fmt.Println(anyType)
}
