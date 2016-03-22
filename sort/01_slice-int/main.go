package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{5, 4, 6, 1, 4, 34, 2}
	// unsorted
	fmt.Println(n)
	sort.Ints(n)
	// slice ordered
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	// reverse order
	fmt.Println(n)

	var x int = 12
	fmt.Println(x)
}
