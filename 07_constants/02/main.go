package main

import "fmt"

//constantes
const (
	_  = iota // 0
	KB = 3 << 2
)

func main() {
	fmt.Println(KB)
	fmt.Print("binary\t\tdecimal\n")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
}
