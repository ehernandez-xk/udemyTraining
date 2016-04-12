/*
no more deadblock, someone is listen whe the go routine is sending.
*/

package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}
