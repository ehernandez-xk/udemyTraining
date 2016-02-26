package main

import "fmt"

// full exrcises https://goo.gl/tGHNhA

//explanation
/*
Write a func foo which can be called in all of these ways
foo(1,2)
foo(1,2,3)
foo(aSlice...)
foo()
*/

func foo(num ...int) {
	fmt.Println("valid ", num)
}

func main() {
	fmt.Println("function foo ")
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3}
	foo(aSlice...)
	foo()
}
