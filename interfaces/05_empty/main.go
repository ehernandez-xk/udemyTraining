package main

import "fmt"

type vehicles interface{}

type car struct {
	ID     string
	whells int
	colors []string
}

type boat struct {
	ID     string
	deep   string
	colors []string
}

func allVehicles(vs []vehicles) {
	fmt.Println(vs)
}

func main() {
	fmt.Println("Vehicles")
	polo := car{"c10", 4, []string{"blue", "black", "red"}}
	BMW := car{"c11", 4, []string{"black", "red"}}
	aaa := boat{"b100", "50mts", []string{"yellow", "brown"}}
	bbb := boat{"b200", "100mts", []string{}}

	all := []vehicles{polo, BMW, aaa, bbb}

	allVehicles(all)

	//criterias := []interface{}{polo, BMW, aaa}

}
