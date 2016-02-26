package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "este es un texto cualquiera, y se hará split en sus palabras"
	scanner := bufio.NewScanner(strings.NewReader(input))

	//Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	//Count the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}
}
