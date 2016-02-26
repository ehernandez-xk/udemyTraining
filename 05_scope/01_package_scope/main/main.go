package main

import (
	"fmt"

	"github.com/ehernandez-xk/udemyTraining/05_scope/01_package_scope/printname"
)

func main() {
	fmt.Println("Testing package variable scope")
	printname.Addname("Eddy")
	printname.Addlastname("Hernandez")
	printname.Fullname()
}
