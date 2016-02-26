package main

import "fmt"

//Passing the parameter as a value
func changeName(name string) {
	name = "pepe"
}

//passing the parameter as a reference
func changeName2(name *string) {
	*name = "pepe"
}

func main() {
	nombre := "eddy"
	changeName(nombre)
	fmt.Println(nombre)
	changeName2(&nombre)
	fmt.Println(nombre)
	//see pass-by-value folder
}
