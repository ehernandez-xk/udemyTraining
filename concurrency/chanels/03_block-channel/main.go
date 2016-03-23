/*
Exercise 14.1: channel_block3.go: Demonstrate the blocking nature of channels by making a
channel, starting a go routine which receives the value from the channel, but only after 15s have
passed, and then after the goroutine putting a value on the channel. Print messages at the different
stages and observe the output.
*/
package main

import (
	"fmt"
	"time"
)

func foo(ch chan int) {
	time.Sleep(15 * 1e9) // 15 seconds
	fmt.Println("received ", <-ch)
}

func main() {
	ch1 := make(chan int)
	go foo(ch1)
	fmt.Println("sending ", 10)
	ch1 <- 10
	fmt.Println("sent ", 10)

}
