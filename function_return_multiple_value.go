package main

import "fmt"

func main() {
	name, success := getName("Rizqi", "Tsani")
	fmt.Println(name, success)
}

func getName(firstName string, lastName string) (string, bool) {
	name := firstName + " " + lastName
	success := true
	return name, success
}
