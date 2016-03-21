package main

// How quickly can we do our multiplication tables?

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func timestable(x int) {
	for i := 0; i < 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, x, x*i)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	for n := 2; n <= 12; n++ {
		go timestable(n)
	}
	wg.Wait()
}
