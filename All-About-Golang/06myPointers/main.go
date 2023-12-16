package main
import (
	"fmt"
)
func main()  {
	fmt.Println("Here is Pointers")

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Actual value pointer",ptr)
	fmt.Println("value of pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println(myNumber)
}