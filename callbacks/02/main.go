package main

import "fmt"

func filter(numbers []int, call func(int) bool) []int {
	xs := []int{}
	//or: var xs []int
	for _, number := range numbers {
		if call(number) {
			xs = append(xs, number)
		}
	}
	return xs
}

func main() {
	xs := filter([]int{0, 1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs) // [2 3 4]
	// callback: passing a func as an argument
}
