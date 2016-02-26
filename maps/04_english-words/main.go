package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// get call under http
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	//Prints status of the request.
	fmt.Println(res.Body)

	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)
	fmt.Println(str)
}
