package main

import (
	"encoding/json"
	"fmt"
)

//Person struct
type Person struct {
	First string
	Last  string `json:"-"`            //don't include last
	Age   int    `json:"wisdom score"` //Change Age
}

func main() {
	p1 := Person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
