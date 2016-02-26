package main

import "fmt"

//Person Create a type person that has fields
type Person struct {
	first string
	last  string
	age   int
}

//DoubleZero struct
type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			first: "pepe",
			last:  "trueno",
			age:   22,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	// output all p1
	fmt.Println(p1)

	fmt.Println(p1.Person.first, p1.Person.last, p1.Person.age)

}
