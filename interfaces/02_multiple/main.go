package main

import (
	"fmt"
	"math"
)

//Square struct
type Square struct {
	side float64
}

//Circle struct
type Circle struct {
	radius float64
}

// square implements shape interface
func (s Square) area() float64 {
	return s.side * s.side
}

// circle implements Shape interface
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//Shape interface, this interface will use to the methods named area() and return float64
// we can say a square is a shape.
// it is a type that defines behavior
type Shape interface {
	area() float64
}

//info receives a kind of shape
func info(z Shape) {
	fmt.Printf("%T area: %v \n", z, z.area())
}

func main() {
	square := Square{10.5}
	circle := Circle{5.5}
	//sends a saquare that is a shape.
	info(square)
	info(circle)

}
