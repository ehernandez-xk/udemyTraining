package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("adding to a chanel")
			c <- i
		}
		close(c) //to close the channel

	}()

	for n := range c {
		fmt.Println(n)
	}
}
