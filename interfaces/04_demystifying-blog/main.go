/*
Blog explanation
http://nathanleclaire.com/blog/2014/07/19/demystifying-golangs-io-dot-reader-and-io-dot-writer-interfaces/
*/

package main

import (
	"fmt"
	"os"
)

//Animal interface
type Animal interface {
	Say() string
	Greet(Animal)
}

//Person struct
type Person struct {
}

//Say Person says
func (p Person) Say() string {
	return "Hey there"
}

//Greet Person greets
func (p Person) Greet(animalToGreet Animal) {
	//defies the relation person to dog
	if _, ok := animalToGreet.(Dog); ok {
		fmt.Println("hello doggie")

	} else {
		// if the person talk with a person
		fmt.Println("Hi")
	}
}

//Dog struct
type Dog struct {
	age   int
	breed string
	owner *Person
}

// Say dog says
func (d Dog) Say() string {
	return "woof"
}

//Growl dog growls
func (d Dog) Growl() {
	fmt.Println("Grrr")
}

//Snuggle dog snuggles
func (d *Dog) Snuggle() {
	fmt.Println("snuggle")
}

//Sniff dog sniffs
func (d Dog) Sniff(animalToSniff Animal) (bool, error) {
	fmt.Println("sniffing")
	return true, nil
}

//Greet dog greets
func (d Dog) Greet(animalToGreet Animal) {
	if _, ok := animalToGreet.(Person); ok {
		d.Snuggle()
	} else {
		friendly, err := d.Sniff(animalToGreet)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error sniffing a non-person")
		}
		if !friendly {
			d.Growl()
		}
	}
}

func main() {
	d1 := Dog{2, "shibe", &Person{}}
	d2 := Dog{3, "poodle", &Person{}}
	p1 := Person{}
	p2 := Person{}

	fmt.Println("--- Dogs say ---")
	d1.Greet(d2)
	d1.Greet(p1)
	fmt.Println(d1.Say())

	fmt.Println("--- People say ---")
	p2.Greet(p1)
	p2.Greet(d1)
	fmt.Println(p2.Say())

}
