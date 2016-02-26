package main

import "fmt"

func myFunc(myParam string) string {
	myParam = "Variable por valor"
	return myParam
}

func printSlice(numbers []string) {
	for i, number := range numbers {
		fmt.Println(i, number)
	}
}

func printMap(numbers map[int]string) {
	numbers[4] = "cuatro"
	for i, number := range numbers {
		fmt.Println(i, number)
	}

}

func main() {
	// function
	myValue := "hello"
	fmt.Print(myFunc(myValue) + "\n")

	//slices
	myNumbers := []string{"uno", "dos", "tres"}
	printSlice(myNumbers)

	//maps
	myMap := map[int]string{1: "uno", 2: "dos", 3: "tres"}
	printMap(myMap)
}
