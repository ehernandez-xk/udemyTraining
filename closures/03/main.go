package main

import "fmt"

func wrapper() func() int {
	var x int
	// the same assigning to 0
	//x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	//outout shows 1 2 and 3
	/* When we call to increment() it saves the value of x in its scope */

}
