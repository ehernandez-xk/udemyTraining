package main

import "fmt"

func main() {
	myMap := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
	}

	fmt.Println(myMap)

	//Validate if exist
	value, exists := myMap[1]
	if exists {
		fmt.Println("value ", value, " exists")
	} else {
		fmt.Println("sorry, the value doesn't exists")
	}

	//Delete if exists
	if val, exists := myMap[1]; exists {
		delete(myMap, 1)
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myMap)
}
