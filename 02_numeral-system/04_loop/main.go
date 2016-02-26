package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		fmt.Print("Decimal \t Binary \t Hexadecimal \t UTF-8 \n")
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
