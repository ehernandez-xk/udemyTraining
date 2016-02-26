package main

import "fmt"

func world() {
	fmt.Println("world!")
}

func hello() {
	fmt.Print("Hello ")
}

func main() {
	defer world()
	// do more things.
	hello()

}
