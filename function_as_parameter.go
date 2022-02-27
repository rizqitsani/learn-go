package main

import "fmt"

func filter(data []int, filterFunction func(int) bool) []int {
	var result []int

	for _, value := range data {
		if isFulfilled := filterFunction(value); isFulfilled {
			result = append(result, value)
		}
	}

	return result
}

func main() {
	myNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	evenNumbers := filter(myNumbers, func(i int) bool { return i%2 == 0 })
	fmt.Println(evenNumbers)

	largerThanFiveNumbers := filter(myNumbers, func(i int) bool { return i > 5 })
	fmt.Println(largerThanFiveNumbers)
}
