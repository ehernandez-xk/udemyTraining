package main

import "fmt"

// prints buzz in number multiples of 5, prints fizz on numbers multiples of 3
// and prints FizzBuzz on numbers multiples on 3 and 5
func main() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, " FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println(i, " Buzz")
		} else if i%3 == 0 {
			fmt.Println(i, " Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
