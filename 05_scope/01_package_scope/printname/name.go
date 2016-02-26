package printname

import "fmt"

var name string

/*
Addname Adds name
*/
func Addname(x string) {
	name = x
	fmt.Println(x)
}
