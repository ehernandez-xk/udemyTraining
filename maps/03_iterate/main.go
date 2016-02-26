package main

import "fmt"

func main() {
	myMap := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
	}

	for key, value := range myMap {
		fmt.Println("key ", key, " value ", value)
	}
}
