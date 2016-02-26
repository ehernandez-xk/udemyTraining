package main

import "fmt"

//This is not common to do it, it is just an example to create your ouwn types

type foo int

func main() {
	var myAge foo
	myAge = 44

	//shows the type and the value
	fmt.Printf("%T %v \n", myAge, myAge)
}
