package main

import "fmt"

//Person struct
type Person struct {
	name string
	last string
	age  int
}

func (p Person) fullName() string {
	return p.name + " " + p.last
}

//Overliad the behave of the println
func (p Person) String() string {
	return "you can print this type"
}

func main() {
	p1 := Person{"pepe", "trueno", 22}
	fmt.Println(p1) // Overlaid the func String()
	fmt.Println(p1.fullName())
	fmt.Println(p1.age)
}
