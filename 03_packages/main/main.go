package main

import (
	"fmt"

	"github.com/ehernandez-xk/udemyTraining/03_packages/stringutil"
)

func main() {
	fmt.Print("printing name " + stringutil.MyName)
	stringutil.Reverse()

}
