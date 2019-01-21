package main
import "fmt"

func givemepear( fruit string) {
	fruit = "pear"
}

func givemeactualpear( fruit *string) {
	*fruit = "pear"
}
func main() {
	fruit := "banana"
	givemepear(fruit)
	fmt.Println(fruit)
	givemeactualpear(&fruit)
	fmt.Println(fruit)

}
