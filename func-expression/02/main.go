package main

import "fmt"

//The makgeGreeter function returns a function that returns a string.?
func makeGreeter() func() string {
	return func() string {
		return "Hello World 2"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Printf("%T \n", greet) //Shows the type, that is a: func() string
	fmt.Println(greet())
}
