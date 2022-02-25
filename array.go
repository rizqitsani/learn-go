package main

import "fmt"

func main() {
	var names [3]string

	fmt.Println(names, len(names))

	names[0] = "Muhammad"
	names[1] = "Rizqi"
	names[2] = "Tsani"

	numbers := [...]int{
		1,
		2,
		3,
	}

	fmt.Println(numbers)
}
