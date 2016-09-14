package main

import (
	"errors"
	"fmt"
)

type valid interface {
	Ok() error
}

type person struct {
	name string
}

func (p person) Ok() error {
	if p.name == "" {
		return errors.New("Person need a name.")
	}
	return nil
}

// do not use empty interface, example purpose
func testPerson(v interface{}) error {
	// Type assertions https://golang.org/ref/spec#Type_assertions
	ob, ok := v.(valid)
	if !ok {
		return errors.New("no metod Ok")
	}
	err := ob.Ok()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// p1 is valid because has name
	p1 := person{name: "eddy"}
	err := testPerson(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Person valid")

	// p1 is not valid, empty name
	p1 = person{name: ""}
	err = testPerson(p1)
	if err != nil {
		fmt.Println(err)
	}

	//p2 does not have Ok method
	p2 := "another type"
	err = testPerson(p2)
	if err != nil {
		fmt.Println(err)
		return
	}

	/*
	   Person valid
	   Person need a name.
	   no metod Ok
	*/
}
