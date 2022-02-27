package main

import (
	"fmt"
	"math"
)

type TwoDimensionalFigure interface {
	getArea() float64
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s Square) getArea() float64 {
	return math.Pow(s.side, 2)
}

func (c Circle) getArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func checkFigure(f TwoDimensionalFigure) {
	switch f.(type) {
	case Square:
		fmt.Println("I'm a square!")
	case Circle:
		fmt.Println("I'm a circle!")
	}
}

func main() {
	var figure TwoDimensionalFigure

	figure = Square{4.0}
	checkFigure(figure)
	fmt.Println(figure.getArea())

	figure = Circle{7.0}
	checkFigure(figure)
	fmt.Println(figure.getArea())
}
