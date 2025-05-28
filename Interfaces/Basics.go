package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	// Create a slice of shapes
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Circle{Radius: 11},
		Rectangle{5, 10},
	}
	for _, shape := range shapes {

		switch shape.(type) {
		case Circle:
			fmt.Println("It's a circle")
		case Rectangle:
			fmt.Println("It's a rectangle")
		}
		printArea(shape)

	}
}
