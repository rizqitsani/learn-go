package main

import "fmt"

func main() {
	switch {
	case 17 == 19:
		{
			fmt.Println("Not printed")
		}
	case 17 == 17:
		{
			fmt.Println("Printed")
		}
	}

	age := 19
	switch age {
	case 17:
		{
			fmt.Println("Young")
		}
	case 19, 20:
		{
			fmt.Println("Either 19 or 20")
		}
	}

	name := "Rizqi"
	switch name {
	case "Muhammad":
		{
			fmt.Println("Hi, Muhammad")
		}
	case "Tsani":
		{
			fmt.Println("Hi, Tsani")
		}
	default:
		{
			fmt.Println("Who are you?")
		}
	}
}
