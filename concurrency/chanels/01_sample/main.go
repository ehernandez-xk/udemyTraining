/*
main() waits for 1 second so that both goroutines can come to completion, if this is not
allowed sendData() doesn’t have the chance to produce its output.
•	 getData() works with an infinite for-loop: this comes to an end when sendData() has
finished and ch is empty.
•	 if we remove one or all go—keywords, the program doesn’t work anymore, the Go runtime
throws a panic:
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	//time.Sleep(1e9)
	for {
		input = <-ch
		fmt.Printf("%s  \n", input)
	}
}

/*
Washington
Tripoli
London
Beijing
Tokio

*/
