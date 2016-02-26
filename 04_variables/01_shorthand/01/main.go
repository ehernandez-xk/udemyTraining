package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	//Shows the value and it value type
	fmt.Printf("%v is an %T \n", a, a)
	fmt.Printf("%v is a %T \n", b, b)
	fmt.Printf("%v is a %T \n", c, c)
	fmt.Printf("%v is a %T \n", d, d)

}
