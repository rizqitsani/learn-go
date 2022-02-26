package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Spy struct {
	Person
	licenseToKill bool
	alias         []string
}

func main() {
	spy := Spy{
		Person: Person{
			name: "Rizqi",
			age:  19,
		},
		licenseToKill: true,
		alias:         []string{"Tsani", "Rizqi Tsani"},
	}

	fmt.Println(spy)
	fmt.Println(spy.name, spy.age, spy.licenseToKill)

	for _, alias := range spy.alias {
		fmt.Println(alias)
	}
}
