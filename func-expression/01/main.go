package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Hello world")
	}
	greeting()

	//get the type
	fmt.Printf("%T \n", greeting)
}
