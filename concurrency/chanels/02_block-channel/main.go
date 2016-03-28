package main

import (
	"fmt"
	"time"
)

/*
The foo and faa runs concurrently

*/

// function sender
func foo(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

//function receiver
func faa(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	ch := make(chan int)

	go foo(ch)
	go faa(ch)

	// If we used these prints (instead of faa function) you will see 0 and 2 in the output,
	// This happens because only 2 items are requested using <-
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	time.Sleep(1e9) // one second to Wait

}
