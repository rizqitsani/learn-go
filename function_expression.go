package main

import "fmt"

func main() {
	foo := func() { fmt.Println("I'm a function expression") }
	foo()
}
