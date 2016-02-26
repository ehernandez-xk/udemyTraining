package printname

import "fmt"

var lastname string

/*
Addlastname Adds last name
*/
func Addlastname(x string) {
	lastname = x
	fmt.Println(x)
}
