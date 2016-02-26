package main

import "fmt"

//print the sum of all the number multiples of 3 and 5
func main() {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%5 == 0 {
			sum = sum + i
		} else if i%3 == 0 {
			sum += i
		}
	}
	fmt.Println("sum are: ", sum)
}
