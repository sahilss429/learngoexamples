package main
import "fmt"

type Counter struct {
	Count int
}

func (c Counter) Increment() {
	c.Count++
}
//c is the local variable inside function and Counter is struct type. 
//c.Count is basically Counter.Count in simple language, or just the clone.
func (c *Counter) IncrementWithPointer() {
	c.Count++
}
//Referencing Counter using a pointer
func main() {
	counter := &Counter{
		Count: 1,
	}
	 //Initalized type struct to 1
	fmt.Println(counter.Count) //Print the value of Count before any change
	counter.Increment()	// 
	fmt.Println(counter.Count)
	counter.IncrementWithPointer()
	fmt.Println(counter.Count)
}
