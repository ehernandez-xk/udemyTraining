package main

import "fmt"

//retuns sum of all parameters
func manyFloats(a ...float64) float64 {
	var sum float64
	for _, number := range a {
		sum += number
	}
	return sum
}

func main() {
	data := []float64{1, 2, 3, 4}
	fmt.Printf("%T \n", data)
	fmt.Println("The sum is: ", manyFloats(data...))

}
