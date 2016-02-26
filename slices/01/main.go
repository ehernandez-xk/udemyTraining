package main

import "fmt"

func main() {
	y := []int{1, 2, 3}
	fmt.Println("cap ", cap(y), "len ", len(y)) //cap  3 len  3
	y = append(y, 4, 5)
	fmt.Println("cap ", cap(y), "len ", len(y)) //cap  6 len  5
}
