//sync.WaitGroup is essentially a counter that we can increase (to indicate we want
//to wait for things), and decrease (to indicate things are done). Then we can tell
//code to wait until the WaitGroup counter reaches zero, which would mean all things have finished.

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func doSomething() {
	fmt.Println("do somethig")
	wg.Done() //This is done
}

func main() {
	fmt.Println("start")
	wg.Add(1) // indicate we are going to wait for one thing
	go doSomething()
	wg.Wait() // here the expected behavior
	fmt.Println("end")
	//wg.Wait() // wait for all things to be done

}
