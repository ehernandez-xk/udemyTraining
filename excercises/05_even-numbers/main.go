package main

import "fmt"

//Even numbers from 0 to 100
func main() {
	for i := 1; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("Even number: ", i)
		}
	}
}
