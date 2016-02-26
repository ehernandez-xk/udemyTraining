package main

import "fmt"

func main() {
	var myMap map[string]string
	//if nil I can't do this:
	//myMap["dos"] = "two"

	var myMap2 = make(map[string]string)
	myMap2["dos"] = "two"
	fmt.Println(myMap, myMap2)
}
