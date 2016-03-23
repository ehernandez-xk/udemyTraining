package main

import "fmt"
import "time"

func foo(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
func faa(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	ch := make(chan int)

	go foo(ch)
	go faa(ch)

	fmt.Println(<-ch)
	time.Sleep(1e9) // one second

}
