package main

import "fmt"

func main() {
	transactions := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		var transaction []int
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
}

// output [[0 1 2 3] [0 1 2 3] [0 1 2 3]]
