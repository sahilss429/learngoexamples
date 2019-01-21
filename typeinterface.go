package main
import "fmt"

type Fruit interface {
	String() string
}

type Apple struct {
	Variety string
}

func (a Apple) String() string {
	return fmt.Sprintf("A %s apple", a.Variety)
}

type Orange struct {
	Size string
}

func (o Orange) String() string {
	return fmt.Sprintf("A %s orange", o.Size )
}
//Created a func with paramter fruit of type Fruit interface.
func PrintFruit(fruit Fruit) {
	fmt.Println("I have this fruit:", fruit.String())
}

func main() {
	//Initialized Apple and Orange struct with value
	apple := Apple{"Golden Delicious"}
	orange := Orange{"large"}
	//Called the function with above initialized structs.
	PrintFruit(apple)
	PrintFruit(orange)
}
