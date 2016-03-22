package main

import (
	"fmt"
	"runtime"
)

var var1 string
var var2 string

//The init funcitons helps to initialize configs to our programs
//it is executed after the variables
// https://golang.org/doc/effective_go.html#init
func init() {
	var1 = "hello" // don't do this
	var2 = "world" // don't do this

	//Defines how many CPUs the program will use.
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Println(var1, " ", var2)

	fmt.Println("my pc has: ", runtime.NumCPU(), " CPUs")

}
