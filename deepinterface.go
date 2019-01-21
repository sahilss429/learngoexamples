package main
import "fmt"

type ProductCatalogue interface {
	All() ([]Product, error)
	Find(string) (Product, error)
}
//Two different implementations of same interface. basically FileProductCatalogue and DBProductCatalogue and two implementations of ProductCatalogue.


type FileProductCatalogue struct {
	//Fields...
}

func (c FileProductCatalogue) All() ([]Product, error) {
	//implementations
}

func (c FileProductCatalogue) Find(name string) (Product, error) {
	//implementations
}

type DBProductCatalogue struct {
	//Fields..
}

func (d DBProductCatalogue) All() ([]Product, error) {
	//implementations..
}

func (d DBProductCatalogue) Find( name string) (Product, error){
	//implementations
}

func LoadProducts(catalogue ProductCatalogue) {
	products, err := catalogue.All()
}

func main() {
	var myCatalogue ProductCatalogue
	myCatalogue = DBProductCatalogue{}
}
