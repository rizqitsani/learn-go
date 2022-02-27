package main

import "fmt"

func bar() (string, int) {
	return "Rizqi", 19
}

func sumVariadic(nums ...int) int {
	var result int

	for _, num := range nums {
		result += num
	}

	return result
}

func sumArray(nums []int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	name, age := bar()
	fmt.Println(name, age)

	myNumbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sumVariadic(myNumbers...))
	fmt.Println(sumArray(myNumbers))
}
