package main

import "fmt"

// This function prints numbers
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

// This is the first goroutine implicit.
func main() {
	// Whe goroutine comes the program automaticaly pass to the next line
	go f(0)

	//We take the output of the background print.
	var input string
	fmt.Scanln(&input)
}
