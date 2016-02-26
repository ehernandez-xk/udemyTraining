package main

import "fmt"

//returns a string with the concatenated parameters
func manyStrings(a ...string) string {
	var allStrings string
	fmt.Println("All parameters: ", a)
	fmt.Printf("The type is: %T \n", a)

	for _, value := range a {
		//fmt.Println(value)
		allStrings += value
	}
	return allStrings
}

func main() {

	//variadic functions
	manyStrings("hola", "mundo", "xxx")

	//variaic arguments
	myStrings := []string{"a", "b", "c", "d"}
	manyStrings(myStrings...)
}
