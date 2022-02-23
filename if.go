package main

import "fmt"

func main() {
	name := "Rizqi"

	if name == "Rizqi" {
		fmt.Println("Hi, Rizqi!")
	} else if name == "Tsani" {
		fmt.Println("Hi, Tsani")
	} else {
		fmt.Println("Who are you?")
	}

	if age := 19; age > 10 {
		fmt.Println("You're old")
	}
}
