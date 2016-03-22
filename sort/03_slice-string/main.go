/*
  use godoc.org/sort to sort the following:
  s := []string{"zeno", "John", "Al", "Jenny"}

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"zeno", "john", "al", "jenny"}
	fmt.Println(s)
	//Two options to sort
	sort.StringSlice(s).Sort()
	//sort.Sort(sort.StringSlice(s))
	fmt.Println(s)

	fmt.Println("len:  ", sort.StringSlice.Len(s))

	fmt.Println("---------------------")
	//sort slice of float64

	f := []float64{20, 2, 10, 3.2}
	fmt.Println(f)
	sort.Float64s(f)
	fmt.Println(f)
	fmt.Println("is sort? ", sort.Float64sAreSorted(f))

	fmt.Println("---------------------")
	// sort slice of int
	n := []int{4, 1, 2, 6, 1, 8}
	fmt.Println(n)
	sort.Ints(n)
	//sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}
