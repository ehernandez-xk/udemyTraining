package main

import "fmt"

// everything in go is pass by value

// a string is passed by value
func changeString(myString string) {
	myString = "new String"
}

// a slice is passed by value but is referenced.
func changeSlice(mySlice []int) {
	mySlice[0] = 22
}

// a map is passed by value but is referenced.
func changeMap(myMap map[string]int) {
	myMap["eddy"] = 11
}

func main() {
	//passing a string
	myString := "hola"
	changeString(myString)
	fmt.Println(myString) // hola

	//passing a slice
	mySlice := []int{1, 2}
	changeSlice(mySlice)
	fmt.Println(mySlice) // [22 2]

	//passing a map
	myMap := map[string]int{
		"eddy": 5,
		"pepe": 12,
	}
	changeMap(myMap)
	fmt.Println(myMap) // map[eddy:11 pepe:12]
}
