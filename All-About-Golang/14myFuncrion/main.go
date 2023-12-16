package main

import (
	"fmt"
)

func Greet()  {
	fmt.Println("Hey budddy whats up????")
}


func adder(val1 int,val2 int) int {
	return val1+ val2
}

func proAdder(values... int) (int, string) {
	total :=0
	for _, val := range values {
		total+=val
	}

	return total , "hey this is your proResult"
}

func main()  {
	Greet()

	result := adder(10,20)
	fmt.Println("result = ",result)

	proresult, proMessage := proAdder(10,20,30,40)

	fmt.Println("Pro Result = ",proresult)
	fmt.Println("Pro Messagae  =", proMessage)
}