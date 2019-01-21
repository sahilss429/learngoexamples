package main
import "fmt"

type Movie struct {
	Actors	[]string
	ReleaseYear	int32
	Rating	float32
	Title	string
}
//An example of Type Method, where movie is variable known as receiver type and Movie is struct.
func (movie Movie) DisplayTitle() string {
	return fmt.Sprintf("%s (%d)", movie.Title, movie.ReleaseYear)
}

func main() {
	Favorite := Movie{
		Title: "DDLJ",
		ReleaseYear: 1991,
		Rating: 5.9,
	}
	Favorite.Actors = []string{
		"ShahRukh Khan",
		"Kajol",
		"Many More",
	}
	fmt.Println("My Favourite movie was", Favorite.Title, "Released on", Favorite.ReleaseYear, "with rating", Favorite.Rating, "and Actors", Favorite.Actors[0], "and", Favorite.Actors[1])

	fmt.Println(Favorite.DisplayTitle())
}
