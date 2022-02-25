package main

import "fmt"

func main() {
	numbers := []int{5, 23, 1, 75, 21}
	fmt.Println(numbers)
	fmt.Println(numbers[:])
	fmt.Println(numbers[2:])
	fmt.Println(numbers[2:4])

	numbers = append(numbers, 2)
	fmt.Println(numbers)

	otherSlice := make([]int, 2)
	otherSlice[0] = 10
	otherSlice[1] = 7

	numbers = append(numbers, otherSlice...)
	fmt.Println(numbers)

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	twoDimesionalSlice := [][]int{numbers, otherSlice}
	fmt.Println(twoDimesionalSlice)

	for _, v1 := range twoDimesionalSlice {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
