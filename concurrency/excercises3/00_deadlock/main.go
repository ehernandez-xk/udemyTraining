/*
fatal error: all goroutines are asleep - deadlock!

*/

package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}
