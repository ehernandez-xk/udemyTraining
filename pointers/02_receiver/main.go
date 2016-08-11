package main

import "fmt"

type car struct {
	color string
}

// It tries to change the color but it is not possible.
// The receiver is passing by value.
func (c car) tryChangeColor() {
	c.color = "yellow"
}

// Changes the color because it passes a pointer of the car.
func (c *car) changeColor() {
	c.color = "red"
}

func main() {
	c1 := car{color: "blue"}
	c1.tryChangeColor()
	fmt.Println(c1.color)
	c1.changeColor()
	fmt.Println(c1.color)
}

/*
output:
 blue
 red
*/
