package main

import "fmt"
import "sync"

import "time"

var wg sync.WaitGroup

func working(i int) {
	fmt.Println(i, " writing code <> ")
	time.Sleep(5 * time.Millisecond)
	wg.Done()
}

func seeFB(i int) {
	fmt.Println(i, " Seeing FB")
	time.Sleep(250 * time.Millisecond)
	wg.Done()
}

func main() {
	//suponiendo que en el día veo el fb 25 veces
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go seeFB(i)
	}

	//suponiendo que miro mi editor  30 veces
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go working(i)
	}
	wg.Wait()
}
