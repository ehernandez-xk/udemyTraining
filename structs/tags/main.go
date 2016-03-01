package main

import (
	"encoding/json"
	"fmt"
)

//Person is a struct of a person
type Person struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MidddleName string `json:"middle_name,omitempty"`
}

func main() {
	jsonstring := `
  {
    "first_name": "eddy",
    "last_name": "hernandez"
  }`

	person := new(Person)
	json.Unmarshal([]byte(jsonstring), person)

	newjson, _ := json.Marshal(person)
	fmt.Printf("%s\n", newjson)
}
