package main

import "fmt"

//Create a type person that has fields
type person struct {
	first string
	last  string
	age   int
}

type DoubleZero

func main() {
	p1 := person{"pepe", "trueno", 22}
	fmt.Println(p1.first, p1.last, p1.age)

}
