package main

import "fmt"

func main() {
	a := 82

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	// shows the reference
	fmt.Println(b)
	//get the value
	fmt.Println(*b)
}
