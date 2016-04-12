/*
	This prints all numbers, I hade to add a sleep to show the numbers.
	go run -race main.go
*/

package main

import "fmt"
import "time"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(1e9)
}
