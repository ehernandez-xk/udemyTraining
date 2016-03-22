package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	raius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.raius * c.raius
}

func info(s shape) {
	fmt.Println("area: ", s.area())
}

func main() {
	c := circle{5}
	info(&c)

	/* we can send the variable location and receive a pointer
	or we can send onlty the variable location like
	info(&c), without the asterisc in the circle. Go knows how to
	figure out. (c circle)
	tip
	Receivers		Values
	(t T)				T and *T
	(t *T)			*T

	*/

}
