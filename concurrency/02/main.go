package main

import (
	"fmt"
	"math/rand"
	"time"
)

// f prints out the numbers from 0 to 10, waiting between 0 and 250 ms after each one.
// The goroutines should now run simultaneously.
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}

	//We take the output of the background print.
	var input string
	fmt.Scanln(&input)
}
