package main

import "fmt"

func main() {
	fmt.Println("runes")

	// prints i
	// string(i) converts to string
	// []byte(string(i)) converts the string to binary

	for i := 50; i < 140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
		// example of when i is equals to 65
		// 65  -  A  -  [65]
	}
	//shows the decimal corresponding to the ASCII table for each rune
	fmt.Println([]byte("eddy"))
	fmt.Println('e') //converts as well
}
