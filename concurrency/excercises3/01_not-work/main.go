/*
go run -race main.go // everything works "correctly" but we have
only 0 in the output
*/

package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	fmt.Println(<-c)

}
