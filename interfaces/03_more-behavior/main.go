package main

import (
	"fmt"
	"math"
)

//Shape interface, defines the shape's behaviors
type Shape interface {
	area() float64
	perimeter() float64
}

//Square struct
type Square struct {
	side float64
}

//Circle struct
type Circle struct {
	radius float64
}

// area of a square
func (s Square) area() float64 {
	return s.side * s.side
}

// perimeter of a square
func (s Square) perimeter() float64 {
	return s.side * 4
}

// area of a circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// perimeter of a circle
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

//info receives a kind of shape
func info(z Shape) {
	fmt.Printf("%T area: %.2f \n", z, z.area())
	fmt.Printf("%T perimeter: %.2f \n", z, z.perimeter())
}

func main() {
	square := Square{10.5}
	circle := Circle{5.5}
	//sends a saquare that is a shape.
	info(square)
	info(circle)

}
