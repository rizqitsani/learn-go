package main

import "fmt"

func main() {
	minScore := 75
	myScore := 90

	fmt.Println(minScore == myScore)
	fmt.Println(minScore != myScore)
	fmt.Println(minScore >= myScore)
	fmt.Println(minScore <= myScore)

	fmt.Println(true && false)
	fmt.Println(true || false)
}
