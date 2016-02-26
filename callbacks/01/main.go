package main

import "fmt"

func myFunc(numbers []int, call func(int)) {
	for _, number := range numbers {
		call(number)
	}
}

func main() {
	fmt.Print("")
	numbers := []int{1, 2, 3, 4}
	myFunc(numbers, func(n int) {
		fmt.Println(n)
	})
	// callback: passing a func as an argument
}
