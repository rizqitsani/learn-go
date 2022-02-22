package main

import "fmt"

var globalName = "Muhammad Rizqi Tsani"

var (
	firstName = "Rizqi"
	lastName  = "Tsani"
)

func main() {
	age := 19
	fmt.Println(age)

	fmt.Println(globalName, firstName, lastName)

	var name string
	fmt.Println(name) // Zero value

	name = "Rizqi Tsani"
	fmt.Println(name)

	name = "Tsani"
	fmt.Println(name)

	var hobby = "code"
	fmt.Println(hobby)
}
