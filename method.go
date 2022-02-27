package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) sayHello() {
	fmt.Printf("I'm %v\n", p.name)
}

func main() {
	p1 := Person{
		"Rizqi",
		18,
	}

	p1.sayHello()
}
