package main

import "fmt"

//Square struct
type Square struct {
	side float64
}

// Adds a method area() to the square struct
func (s Square) area() float64 {
	return s.side * s.side
}

//Shape interface, this interface will use to the methods named area() and return float64
// we can say a square is a shape.
type Shape interface {
	area() float64
}

//info receives a kind of shape
func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	shape := Square{10.5}
	//sends a saquare that is a shape.
	info(shape)
}
