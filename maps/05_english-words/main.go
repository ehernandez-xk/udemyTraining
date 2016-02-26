package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// get call under http
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		// no definition about the words, that we use ""
		words[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}
	i := 0
	// no Blanck identifier is needed, we are using only the key and not the value
	for k := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}
}
