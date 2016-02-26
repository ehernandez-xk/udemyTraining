package main

import "fmt"

func main() {
	fmt.Println('a')
	foo := 'a'                     //Get the rune
	fmt.Println(foo)               // shows Decimal 97
	fmt.Printf("%T \n", foo)       //shows the type int32
	fmt.Printf("%v \n", rune(foo)) //shows Decimal 97
	fmt.Printf("%T \n", rune(foo)) //shows the type int32
}
