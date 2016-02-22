package main

import "fmt"

func main() {
	// Examples decimal, binary and hexadecimal
	// https://godoc.org/fmt
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#x \n", 42, 42, 42)
	fmt.Printf("soy %b \n", 10)
	/* output
	42 - 101010 - 2a
	42 - 101010 - 0x2a
	42 - 101010 - 0X2A
	42 	 101010 	 0x2a
	*/
}
