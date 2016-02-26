package main

import "fmt"

//single function
func singleFunc(name string) {
	fmt.Println(name)
}

// return string length
func funcAndReturn(name string) int {
	return len(name)
}

//return with name assigned
func returnAssigned(name string) (s string) {
	s = "hola " + name
	return
}

//multiple returns
func multipleReturns(name string, age int) (string, int) {
	return "hi " + name, age
}

func main() {

	singleFunc("eddy")
	fmt.Println(funcAndReturn("Eddy"))
	fmt.Println(returnAssigned("eddy"))
	fmt.Println(multipleReturns("Eddy", 33))
}
