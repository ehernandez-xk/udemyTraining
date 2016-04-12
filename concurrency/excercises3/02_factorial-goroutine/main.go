/*
use this example and re-write it using go routines and channles.
*/
package main

import "fmt"

func main() {
	ch := factorial(4)
	for num := range ch {
		fmt.Println("total ", num)
	}
}

func factorial(num int) chan int {
	ch := make(chan int)
	go func() {
		total := 1
		for i := num; i > 0; i-- {
			total *= i
		}
		ch <- total
		close(ch) //this prevent deadlock. to know when finish
	}()
	return ch
}
