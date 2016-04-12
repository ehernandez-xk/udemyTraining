/*
	it works, note: important to have latest version of go 1.5.2 or greater
	I closed the channel to know when the channel finish.
	go run -race main.go
*/

package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
