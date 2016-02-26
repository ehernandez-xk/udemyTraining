package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println("for 1: ", i)
	}

	i := 0
	for i < 5 {
		fmt.Println("for 2: ", i)
		i++
	}
}
