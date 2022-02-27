package main

import "fmt"

func main() {
	name := getFullName("Rizqi", "Tsani")
	fmt.Println(name)
}

func getFullName(firstName string, lastName string) string {
	return firstName + " " + lastName
}
