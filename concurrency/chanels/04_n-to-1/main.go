// many function writing to a channel
// go run -race main.go

package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2) // best way to prevent race condition

	go func() {
		// wg.Add(1) // generates race condition
		for i := 0; i < 10; i++ {
			fmt.Println("adding value 1")
			c <- i
		}
		wg.Done()
	}()

	go func() {
		// wg.Add(1) // generates race condition
		for i := 10; i < 20; i++ {
			fmt.Println("adding value 2")
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
