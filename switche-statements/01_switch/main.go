package main

import "fmt"

func main() {
	switch "cuatro" {
	case "uno":
		fmt.Println("uno")
	case "dos":
		fmt.Println("dos")
		fallthrough
	case "tres":
		fmt.Println("tres")
	case "cuatro", "cinco":
		fmt.Println("cuatro o cinco")
	default:
		fmt.Println("saber")
	}
}
