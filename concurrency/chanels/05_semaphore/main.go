/*
This example show the same output like 04_n-to-1
The difference is not using WaitGroup
*/

package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 10; i < 20; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done   // is waiting to the done, pass to the next line
		<-done   // waiting another value
		close(c) // when two values of "done" have finished
	}()

	for n := range c {
		fmt.Println(n)
	}
}
