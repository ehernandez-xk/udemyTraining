package main

import "fmt"

//Movie Creates the movie struct
type Movie struct {
	Authors []string
	Rating  int
	Title   string
	year    int
	otra    string
}

// Adds a movie method
func (movie Movie) displayTitle() string {
	fmt.Println("The Title is: ", movie.Title)
	return movie.Title
}

func main() {
	fmt.Println("structures")
	// new movie struct and assigns values
	myMovie := Movie{
		Authors: []string{"Pepe", "Artur"},
		Title:   "La Peli",
		Rating:  9,
		year:    2010,
	}
	//assigns values
	myMovie.year = 2000
	//to access
	fmt.Println(myMovie)
	fmt.Println(myMovie.Authors)

	//using the method of the struct
	myMovie.displayTitle()

}
