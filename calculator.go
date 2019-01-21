package main
import "fmt"

func add(a , b int) string {
	return fmt.Sprintf("sum is %d", a + b )
}

func TakeInput() int {
	var a int
	fmt.Println("Enter a number")
	_, err := fmt.Scanf("%d", &a)
		if err != nil {
			fmt.Println("something happened")
		}
	return a
}

func main() {
	a := TakeInput()
	b := TakeInput()
	fmt.Println(add(a,b))
}
