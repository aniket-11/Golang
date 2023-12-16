package main 
import(
	"fmt"
)

func main()  {
	fmt.Println("Here is Array Awww")

	//declaration type 1
	var fruitList [5]string

	fruitList[0] = "Apple"
	fruitList[1] = "Grapes"
	fruitList[2] = "Kiwi"
	fruitList[3] = "Orange"

	
	for i := 0; i < 4; i++ {
		fmt.Println("fruitName = ",fruitList[i])
	}
	fmt.Println(fruitList)
	fmt.Println("Fruit list is ", len(fruitList))

	var vegList = [3]string{"potato","beans","tomato"}
	fmt.Println("vegList =",vegList)

}