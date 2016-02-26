package main

import "fmt"

//single const
const my1 string = "const1"

//multi const
const (
	Pi       = 3.1416
	Language = "Go"
)

//iota example
const (
	A = iota
	B
	C
)

func main() {
	const my2, my3 string = "hola2", "hola3"
	fmt.Println(my1, my2, my3)
	fmt.Println(Pi, Language)
	fmt.Println("iota ", A, B, C)
}
