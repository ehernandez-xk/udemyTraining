/*
use this example and re-write it using go routines and channles.
*/
package main

import "fmt"

func main() {
	num := factorial(4)
	fmt.Println("total ", num)
}

func factorial(num int) int {
	total := 1
	for i := num; i > 0; i-- {
		total *= i
	}
	return total
}
