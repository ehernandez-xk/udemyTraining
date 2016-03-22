/*
examples from godoc.org/sort
*/
package main

import (
	"fmt"
	"sort"
)

// Person struct
type Person struct {
	Name string
	Age  int
}

// Overides Println
func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// People implements sort.Interface for []Person based on the Age field
type People []Person

//attach methods to People
func (a People) Len() int           { return len(a) }
func (a People) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a People) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	friends := []Person{
		{"bob", 31},
		{"john", 42},
		{"michael", 17},
		{"jenny", 26},
	}
	fmt.Println(friends)
	sort.Sort(People(friends))
	fmt.Println(friends)
}
