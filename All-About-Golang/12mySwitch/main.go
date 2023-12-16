package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("Here is my switch game ")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is ",diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 You may open")
	case 2:
		fmt.Println("Youe may move dice to spot 2")	
	case 3:
		fmt.Println("Youe may move dice to spot 3")	
	case 4:
		fmt.Println("Youe may move dice to spot 4")	
	case 5:
		fmt.Println("Youe may move dice to spot 5")	
	case 6:
		fmt.Println("Youe may move dice to spot 6 and roll dice again")	
	default:
		fmt.Println("What was that")		
	}
}