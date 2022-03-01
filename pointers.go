package main

import "fmt"

type Citizen struct {
	name    string
	country string
}

func changeCountry(citizen *Citizen) {
	citizen.country = "Indonesia"
}

func main() {
	c1 := Citizen{
		name:    "Rizqi",
		country: "Singapore",
	}
	fmt.Println(c1)

	changeCountry(&c1)
	fmt.Println(c1)
}
