package main

//encode decode

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//Person struct
type Person struct {
	First      string
	Last       string
	Age        int
	noExported int
}

func main() {

	/*  ENCODE */
	fmt.Println("----------", "ENCODE")
	p1 := Person{"James", "Bond", 20, 007}
	// we write out to the stdout
	//json.NewEncoder(os.Stdout).Encode(p1)

	// step by step
	enco := json.NewEncoder(os.Stdout)
	fmt.Printf("%T \n", enco)
	enco.Encode(p1)

	/*  DECODE */
	fmt.Println("----------", "DECODE")
	var p2 Person
	rdr := strings.NewReader(`{"First":"James", "Last":"Bond", "Age":20}`)
	//json.NewDecoder(rdr).Decode(&p2)

	deco := json.NewDecoder(rdr)
	deco.Decode(&p2)
	fmt.Println("My struct: ", p2)
}
