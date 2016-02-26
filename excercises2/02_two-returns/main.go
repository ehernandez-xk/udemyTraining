package main

import "fmt"

// full exrcises https://goo.gl/tGHNhA

//explanation
/*
Modify the previous exercise to use func expresion
*/
func main() {
	myFunc := func(number int) (int, bool) {
		return number / 2, number%2 == 0
	}
	fmt.Println(myFunc(1))
	fmt.Println(myFunc(2))
	fmt.Println(myFunc(10))

}
