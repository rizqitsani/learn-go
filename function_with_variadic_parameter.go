package main

import "fmt"

func main() {
	total := sumAllNumbers(1, 2, 3, 4, 5)
	fmt.Println(total)
}

func sumAllNumbers(nums ...int) int {
	fmt.Printf("%T\n", nums)
	result := 0

	for _, value := range nums {
		result += value
	}

	return result
}
