package main

import "fmt"
import "encoding/json"

type person struct {
	First      string
	Last       string
	Age        int
	noExported bool //Marshal can't access to this field
}

func main() {
	p1 := person{
		First:      "pepe",
		Last:       "trueno",
		Age:        22,
		noExported: false,
	}
	fmt.Println("object: ", p1)
	bs, _ := json.Marshal(p1) // convert to slices of bytes
	fmt.Println("slice of bytes: ", bs)
	fmt.Printf("type: %T \n", bs)
	fmt.Println("string: ", string(bs))

}
