package main
import (
	"fmt"
	"sort"
	
)
 
func main()  {
	fmt.Println("Here is the Slice Data Structure")
	var fruitList = []string{"Orange","Apple"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango","GreenApple")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	fruitList = append(fruitList[:2])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 135
	highScores[1] = 235
	highScores[2] = 335
	highScores[3] = 435
	// highScores[4] = 435 it is not working bacuase of capcity of value index is 4


	highScores = append(highScores, 55, 60 , 65) 

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)


	var courses = []string{"ds","ml","stqa","wt","react"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	

}