package main

import (
	"encoding/json"
	"fmt"
)

//Person struct
type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person
	fmt.Println(p1)

	bs := []byte(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.Unmarshal(bs, &p1) //*

	fmt.Println("----------")
	fmt.Println(p1)

	//remember Go passes the paramenters as a value.
	//I'm including the pointer so the values are changed.
}
