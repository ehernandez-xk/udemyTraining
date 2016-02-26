package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the a book moby dick
	res, err := http.Get("http://gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}

	//scan page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// Set the split function for the scanning operations
	scanner.Split(bufio.ScanWords)

	// Create slice to hold counts
	buckets := make([]int, 200)

	//Loop over the words, according to the first letter of the word, I increment
	// the value of every match word. at the end I will have a slice with the count of
	// each word.

	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}
	// We have a slice with and index from 0 to 199

	//fmt.Println(buckets[65:122])

	// showing how many words begins with letter A
	fmt.Println("words begining with A ", buckets[65])

	for index, value := range buckets {
		fmt.Println("index ", index, " letter ", string(index), " value ", value)
	}
}

// HashBucket receives a string get the first letter of the word and converts to int
func HashBucket(word string) int {
	return int(word[0])
}
