package main

import "fmt"

// full exrcises https://goo.gl/tGHNhA

//explanation
/*
create a funtion that takes 1 integer and returns two values
the first should be the argument divided by 2
the second should be a boolean that let's us know whether or not the
argumente was even
example
myFunc(1) returns (0, false)
myFunc(2) returns (1, true)

*/

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func myFunc(number int) (int, bool) {
	return number / 2, isEven(number)
	// simple way is
	//return number /2, number%2 == 0
}

func main() {
	fmt.Println(myFunc(1))
	fmt.Println(myFunc(2))
	fmt.Println(myFunc(10))

}
