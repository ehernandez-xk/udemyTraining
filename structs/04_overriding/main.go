package main

import "fmt"

//Person struct
type Person struct {
	first string
	last  string
	age   int
}

//DoubleZero struct
type DoubleZero struct {
	Person
	LicenseTokill bool
}

func (p Person) greeting() string {
	return "Hello Person"
}

func (d DoubleZero) greeting() string {
	return "Hello DoubleZero"
}

func main() {
	p1 := Person{
		first: "pepe",
		last:  "trueno",
		age:   20,
	}
	fmt.Println(p1)            //{pepe trueno 20}
	fmt.Println(p1.greeting()) //Hello Person

	p2 := DoubleZero{
		Person: Person{
			first: "paco",
			last:  "reyes",
			age:   22,
		},
		LicenseTokill: true,
	}
	fmt.Println(p2)                   //{{paco reyes 22} true}
	fmt.Println(p2.greeting())        //Hello DoubleZero
	fmt.Println(p2.Person.greeting()) //Hello Person
}
