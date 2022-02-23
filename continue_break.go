package main

import "fmt"

func main() {
	x := 1
	for {
		if x > 20 {
			break
		}

		if x%2 == 0 {
			x++
			continue
		}

		fmt.Println(x)
		x++
	}
}
