package main 
import (
	"fmt"
)

func main()  {
	fmt.Println("Here is an loops")

	days := []string{"Friday","Monday","Tuesday","Wednesday","Thursday"}

	fmt.Println(days)


	//Method1

	for d:=0 ; d < len(days); d++ {
		fmt.Println(days[d])
	}

	//Method2
	for i := range days{
		fmt.Println(days[i])
	}

	//Method3

	for index,day := range days {
		fmt.Printf("the index %v having a day is %v\n",index,day)
	}

	//Method 4 comma okay syntax

	for _,day :=range days{
		fmt.Printf("Today is %v\n",day)
	}

	//Method 5 likewise while

	rougueValue:=1
	for rougueValue < 10 {

		if rougueValue==2 {
			goto lco
		}

		if rougueValue == 5{
			rougueValue++
			continue
		}
		fmt.Println("Value= ",rougueValue)
		
		if rougueValue == 7{
			break
		}
		rougueValue++
	}

	//goto statement 

	lco:
		fmt.Println("jumoing at aniket.com")
	

}