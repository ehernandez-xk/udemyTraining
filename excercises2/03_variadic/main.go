package main

import "fmt"

// full exrcises https://goo.gl/tGHNhA

//explanation
/*
Write a func with one variadic parameter that finds the greatest number in a list of numbers
*/

func myFunc(numbers ...int) int {
	var greatest int
	for _, num := range numbers {
		if num > greatest {
			greatest = num
		}
	}
	return greatest
}

func main() {
	fmt.Println("greatest number is: ", myFunc(1, 10, 2, 21, 8, 3, 9, 4))
}
