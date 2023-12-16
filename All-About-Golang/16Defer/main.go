package main
import (
	"fmt"
)
func main()  {
	defer fmt.Println("phase1")
	defer fmt.Println("phase2")
	defer fmt.Println("phase3")
	fmt.Println("phase")
	myDefer()
}

func myDefer()  {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
